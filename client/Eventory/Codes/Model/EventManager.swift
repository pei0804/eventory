//
//  EventManager.swift
//  Eventory
//
//  Created by jumpei on 2016/08/19.
//  Copyright © 2016年 jumpei. All rights reserved.
//

import Alamofire
import ObjectMapper
import RealmSwift
import SwiftTask

class EventManager {
    
    static let sharedInstance = EventManager()
    
    private init() {}
    
    private var realm: Realm {
        
        guard let realm = try? Realm() else {
            fatalError("Realm error")
        }
        return realm
    }
    
    func getSelectNewEventAll() -> [EventSummary] {
        
        let genres: [String]! = UserRegister.sharedInstance.getUserSettingGenres()
        let places: [String]! = UserRegister.sharedInstance.getUserSettingPlaces()
        var selectGenre: String = ""
        
        // クエリ文字の結合が1回目かチェックする
        var firstFlg: Bool = false
        for genre in genres {
            if !firstFlg {
                selectGenre += "AND (title CONTAINS[c] '\(genre)' "
                firstFlg = true
            }
            selectGenre += "OR title CONTAINS[c] '\(genre)' "
        }
        // 0は許容していないがバグ回避のためチェック
        if genres.count != 0 {
            selectGenre += ")"
        }
        
        firstFlg = false
        for place in places {
            if !firstFlg {
                selectGenre += "AND (address CONTAINS[c] '\(place)' OR  place CONTAINS[c] '\(place)'"
                firstFlg = true
            }
            selectGenre += "OR address CONTAINS[c] '\(place)' OR place CONTAINS[c] '\(place)' "
        }
        // 0は許容していないがバグ回避のためチェック
        if places.count != 0 {
            selectGenre += ")"
        }
        
        let events: Results<Event> = self.realm.objects(Event).filter("checkStatus == \(CheckStatus.NoCheck.rawValue) \(selectGenre)").sorted("stratAt")
        return setEventInfo(events)
    }
    
    func getNewEventAll(term: String) -> [EventSummary] {
        
        let termArr: [String] = term.componentsSeparatedByCharactersInSet(NSCharacterSet.whitespaceCharacterSet())
        var selectGenre: String = ""
        
        var firstFlg: Bool = false
        
        if term != "" {
            for termWord in termArr {
                if !firstFlg {
                    selectGenre += "AND (title CONTAINS[c] '\(termWord)' OR address CONTAINS[c] '\(termWord)' OR  place CONTAINS[c] '\(termWord)' "
                    firstFlg = true
                } else {
                    selectGenre += "OR title CONTAINS[c] '\(termWord)' OR address CONTAINS[c] '\(termWord)' OR  place CONTAINS[c] '\(termWord)' "
                }
            }
            if termArr.count != 0 {
                selectGenre += ")"
            }
        }
        
        let events: Results<Event> = self.realm.objects(Event).filter("checkStatus == \(CheckStatus.NoCheck.rawValue) \(selectGenre)").sorted("stratAt")
        return setEventInfo(events)
    }
    
    func getKeepEventAll() -> [EventSummary] {
        
        let events: Results<Event> = self.realm.objects(Event).filter("checkStatus == \(CheckStatus.Keep.rawValue)").sorted("stratAt")
        return setEventInfo(events)
    }
    
    func getNoKeepEventAll() -> [EventSummary] {
        
        let events: Results<Event> = self.realm.objects(Event).filter("checkStatus == \(CheckStatus.NoKeep.rawValue)").sorted("stratAt")
        return setEventInfo(events)
    }
    
    func setEventInfo(searchEvents: Results<Event>) -> [EventSummary] {
        
        var eventSummaries: [EventSummary] = [EventSummary]()
        for event in searchEvents {
            let eventSummary: EventSummary = EventSummary()
            eventSummary.id         = event.id
            eventSummary.apiId      = event.apiId
            eventSummary.eventId    = event.eventId
            eventSummary.title      = event.title
            //eventSummary.desc = event.desc
            eventSummary.url        = event.url
            eventSummary.limit      = event.limit
            eventSummary.accepted   = event.accepted
            //eventSummary.waitlisted = event.waitlisted
            eventSummary.address    = event.address
            eventSummary.place      = event.place
            eventSummary.stratAt    = event.stratAt
            eventSummary.endAt      = event.endAt
            eventSummary.checkStatus = event.checkStatus
            eventSummaries.append(eventSummary)
        }
        return eventSummaries
    }
    
    func eventInitializer() {
        
        guard let event: Results<Event> = self.realm.objects(Event) else {
            return
        }
        if event.count == 0 {
            self.fetchNewEvent()
        } else {
            self.realm.beginWrite()
            let oldLocations = self.realm.objects(Event).filter(NSPredicate(format:"stratAt < %@", NSDate().dateByAddingTimeInterval(-86400)))
            EventManager.sharedInstance.fetchNewEvent()
            self.realm.delete(oldLocations)
            do {
                try realm.commitWrite()
            }
            catch {
                fatalError("Realm can not delete")
            }
        }
    }
    
    func keepAction(id: Int, isKeep: Bool) {
        
        if let thisEvent = self.realm.objects(Event).filter("id == \(id)").first {
            if isKeep {
                
                do {
                    try self.realm.write {
                        thisEvent.checkStatus = CheckStatus.Keep.rawValue
                    }
                }
                catch {
                    fatalError("Realm can not wirte")
                }
                
            } else {
                
                do {
                    try self.realm.write {
                        thisEvent.checkStatus = CheckStatus.NoKeep.rawValue
                    }
                }
                catch {
                    fatalError("Realm can not wirte")
                }
            }
        }
    }
    
    func fetchNewEvent() -> Task<Float, String, NSError?> {
        
        var updatedAt = UserRegister.sharedInstance.getUserEventInfoUpdateTime()
        let result = self.realm.objects(Event)

        // 初回ダウンロードは更新時間は関係なしに取ってくるが、
        // サーバーとの通信エラーが発生するとそのまま更新時間がセットされてしまうので、
        // それの対策のために便宜的にupdatedAtを空にしている。
        if(result.count == 0) {
            updatedAt = ""
        }

        return Task<Float, String, NSError?> { progress, fulfill, reject, configure in
            Alamofire.request(.GET, "https://eventory-155000.appspot.com/api/smt/events", parameters: ["updated_at": updatedAt]).responseJSON { response in
                //            Alamofire.request(.GET, "http://ganbaruman.xyz:8080/api/smt/events").responseJSON { response in
                guard let json = response.result.value as? Array<Dictionary<String,AnyObject>> else {
                    reject(nil)
                    return
                }
                json.forEach {
                    do {
                        guard let event = Mapper<Event>().map($0) else {
                            return
                        }
                        try self.realm.write {
                            self.realm.create(Event.self,
                                value:[
                                    "id":event.id,
                                    "eventId"   : event.eventId,
                                    "apiId"     : event.apiId,
                                    "title"     : event.title,
                                    //"desc" : event.desc,
                                    "url"       : event.url,
                                    "limit"     : event.limit,
                                    "accepted"  : event.accepted,
                                    //"waitlisted" : event.waitlisted,
                                    "address"   : event.address,
                                    "place"     : event.place,
                                    "stratAt"   : event.stratAt,
                                    "endAt"     : event.endAt
                                ],
                                update: true)
                        }
                    } catch {}
                }
                fulfill("SUCCESS")
                UserRegister.sharedInstance.setUserEventInfoUpdateTime()
            }
        }
    }
    
    func genreInitializer() ->  [Dictionary<String, AnyObject>] {
        
        let genreArray: [Dictionary<String, AnyObject>] = [
            [
                "name": "PHP",
                "status": false
            ],
            [
                "name": "Javascript",
                "status": false
            ],
            [
                "name": "Java",
                "status": false
            ],
            [
                "name": "Swift",
                "status": false
            ],
            [
                "name": "LT",
                "status": false
            ],
            [
                "name": "HTML",
                "status": false
            ],
            [
                "name": "CSS",
                "status": false
            ],
            [
                "name": "JQuery",
                "status": false
            ]
        ]
        return genreArray
    }
    
    
    func placesInitializer() ->  [Dictionary<String, AnyObject>] {
        
        let place: [Dictionary<String, AnyObject>] = [
            [
                "name": "北海道",
                "status": false
            ],
            [
                "name": "青森県",
                "status": false
            ],
            [
                "name": "岩手県",
                "status": false
            ],
            [
                "name": "宮城県",
                "status": false
            ],
            [
                "name": "秋田県",
                "status": false
            ],
            [
                "name": "山形県",
                "status": false
            ],
            [
                "name": "福島県",
                "status": false
            ],
            [
                "name": "茨城県",
                "status": false
            ],
            [
                "name": "栃木県",
                "status": false
            ],
            [
                "name": "群馬県",
                "status": false
            ],
            [
                "name": "埼玉県",
                "status": false
            ],
            [
                "name": "千葉県",
                "status": false
            ],
            [
                "name": "東京都",
                "status": false
            ],
            [
                "name": "神奈川県",
                "status": false
            ],
            [
                "name": "新潟県",
                "status": false
            ],
            [
                "name": "富山県",
                "status": false
            ],
            [
                "name": "石川県",
                "status": false
            ],
            [
                "name": "福井県",
                "status": false
            ],
            [
                "name": "山梨県",
                "status": false
            ],
            [
                "name": "岐阜県",
                "status": false
            ],
            [
                "name": "静岡県",
                "status": false
            ],
            [
                "name": "愛知県",
                "status": false
            ],
            [
                "name": "三重県",
                "status": false
            ],
            [
                "name": "滋賀県",
                "status": false
            ],
            [
                "name": "京都府",
                "status": false
            ],
            [
                "name": "大阪府",
                "status": false
            ],
            [
                "name": "兵庫県",
                "status": false
            ],
            [
                "name": "奈良県",
                "status": false
            ],
            [
                "name": "和歌山県",
                "status": false
            ],
            [
                "name": "鳥取県",
                "status": false
            ],
            [
                "name": "島根県",
                "status": false
            ],
            [
                "name": "岡山県",
                "status": false
            ],
            [
                "name": "広島県",
                "status": false
            ],
            [
                "name": "山口県",
                "status": false
            ],
            [
                "name": "徳島県",
                "status": false
            ],
            [
                "name": "香川県",
                "status": false
            ],
            [
                "name": "愛媛県",
                "status": false
            ],
            [
                "name": "高知県",
                "status": false
            ],
            [
                "name": "福岡県",
                "status": false
            ],
            [
                "name": "佐賀県",
                "status": false
            ],
            [
                "name": "長崎県",
                "status": false
            ],
            [
                "name": "熊本県",
                "status": false
            ],
            [
                "name": "大分県",
                "status": false
            ],
            [
                "name": "宮崎県",
                "status": false
            ],
            [
                "name": "鹿児島県",
                "status": false
            ],
            [
                "name": "沖縄県",
                "status": false
            ]
        ]
        return place
    }
}
