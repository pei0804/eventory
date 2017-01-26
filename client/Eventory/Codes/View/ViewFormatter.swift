//
//  ViewFormatter.swift
//  Eventory
//
//  Created by jumpei on 2016/08/21.
//  Copyright © 2016年 jumpei. All rights reserved.
//

import Foundation

class ViewFormaatter {
    static let sharedInstance = ViewFormaatter()
    
    private init() {}
    
    func setEventDate(eventSummary: EventSummary) -> String {
        let dateFormatter = NSDateFormatter()
        dateFormatter.locale = NSLocale(localeIdentifier: "ja_JP")
        dateFormatter.dateFormat = "yyyy年MM月dd日 HH:mm"
        let startDate: String = dateFormatter.stringFromDate(eventSummary.stratAt)
        dateFormatter.dateFormat = "MM月dd日 HH:mm"
        let endDate: String = dateFormatter.stringFromDate(eventSummary.endAt)
        return "\(startDate)〜\(endDate)"
    }
}
