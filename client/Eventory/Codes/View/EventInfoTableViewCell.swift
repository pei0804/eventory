//
//  EventInfoTableViewCell.swift
//  Eventory
//
//  Created by jumpei on 2016/08/19.
//  Copyright © 2016年 jumpei. All rights reserved.
//

import UIKit

class EventInfoTableViewCell: UITableViewCell {
    
    override func awakeFromNib() {
        super.awakeFromNib()
        self.separatorInset = UIEdgeInsetsZero
        self.layoutMargins = UIEdgeInsetsZero
    }
    
    override func setSelected(selected: Bool, animated: Bool) {
        super.setSelected(selected, animated: animated)
    }
    
    @IBOutlet weak var noKeepButton: NoKeepButton!
    @IBOutlet weak var keepButton: KeepButton!
    @IBOutlet weak var dateLbl: UILabel!
    @IBOutlet weak var apiNameLbl: UILabel!
    @IBOutlet weak var eventStatusLbl: UILabel!
    @IBOutlet weak var addressLbl: UILabel!
    @IBOutlet weak var titileLbl: UILabel!
    @IBOutlet weak var titleLblWidth: NSLayoutConstraint!
    
    var id: Int = 0
    var indexPath: NSIndexPath = NSIndexPath(index: 0)
    var eventSummary: EventSummary = EventSummary()
    let apiFromMessage = "情報提供元："
    
    private func keepEvent() {
        
        self.titileLbl.textColor = Colors.main
        self.keepButton.active()
        self.noKeepButton.noActive()
    }
    
    private func noKeepEvent() {
        
        self.titileLbl.textColor = Colors.noKeep
        self.noKeepButton.active()
        self.keepButton.noActive()
    }
    
    private func noCheckEvent() {
        
        self.titileLbl.textColor = Colors.noCheck
        self.noKeepButton.noActive()
        self.keepButton.noActive()
    }
    
    @IBAction func keepButton(sender: AnyObject) {
        
        EventManager.sharedInstance.keepAction(id, isKeep: true)
        self.eventSummary.checkStatus = CheckStatus.Keep.rawValue
        self.keepEvent()
    }
    
    @IBAction func noKeepButton(sender: AnyObject) {
        
        EventManager.sharedInstance.keepAction(id, isKeep: false)
        self.eventSummary.checkStatus = CheckStatus.NoKeep.rawValue
        self.noKeepEvent()
    }
    
    func bind(eventSummary: EventSummary, viewPageClass: CheckStatus, indexPath: NSIndexPath) {
        
        self.eventSummary = eventSummary
        
        // 開いている画面
        if CheckStatus.Keep.rawValue == viewPageClass.rawValue {
            self.keepEvent()
        } else if CheckStatus.NoKeep.rawValue == viewPageClass.rawValue {
            self.noKeepEvent()
        }
        
        // イベントの情報
        if self.eventSummary.checkStatus == CheckStatus.NoCheck.rawValue {
            self.noCheckEvent()
        } else if self.eventSummary.checkStatus == CheckStatus.Keep.rawValue {
            self.keepEvent()
        } else if self.eventSummary.checkStatus == CheckStatus.NoKeep.rawValue {
            self.noKeepEvent()
        }
        
        if eventSummary.apiId == ApiId.Atdn.rawValue {
            self.apiNameLbl.text = apiFromMessage + ApiId.Atdn.getName()
            self.eventStatusLbl.text = "\(eventSummary.accepted) / 定員\(eventSummary.limit)人"
        } else if eventSummary.apiId == ApiId.Connpass.rawValue {
            self.apiNameLbl.text = apiFromMessage + ApiId.Connpass.getName()
            self.eventStatusLbl.text = "定員\(eventSummary.limit)人"
        } else if eventSummary.apiId == ApiId.Doorkeeper.rawValue {
            self.eventStatusLbl.text = "\(eventSummary.accepted) / 定員\(eventSummary.limit)人"
            self.apiNameLbl.text = apiFromMessage + ApiId.Doorkeeper.getName()
        }
        
        self.indexPath = indexPath
        
        self.titileLbl.text = eventSummary.title
        self.titileLbl.numberOfLines = 0
        self.titileLbl.lineBreakMode = .ByWordWrapping
        
        self.addressLbl.text = eventSummary.address != "" ? eventSummary.address : "開催地未定"
        self.eventSummary.eventDate = ViewFormaatter.sharedInstance.setEventDate(eventSummary)
        self.dateLbl.text = eventSummary.eventDate
        
        self.id = eventSummary.id
    }
    
}
