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
        separatorInset = UIEdgeInsetsZero
        layoutMargins = UIEdgeInsetsZero
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
        titileLbl.textColor = Colors.main
        keepButton.active()
        noKeepButton.noActive()
    }
    
    private func noKeepEvent() {
        titileLbl.textColor = Colors.noKeep
        noKeepButton.active()
        keepButton.noActive()
    }
    
    private func noCheckEvent() {
        titileLbl.textColor = Colors.noCheck
        noKeepButton.noActive()
        keepButton.noActive()
    }
    
    @IBAction func keepButton(sender: AnyObject) {
        EventManager.sharedInstance.keepAction(id, isKeep: true)
        eventSummary.checkStatus = CheckStatus.Keep.rawValue
        keepEvent()
    }
    
    @IBAction func noKeepButton(sender: AnyObject) {
        EventManager.sharedInstance.keepAction(id, isKeep: false)
        eventSummary.checkStatus = CheckStatus.NoKeep.rawValue
        noKeepEvent()
    }
    
    func bind(eventSummary: EventSummary, viewPageClass: CheckStatus, indexPath: NSIndexPath) {
        
        self.eventSummary = eventSummary
        
        // 開いている画面
        if CheckStatus.Keep.rawValue == viewPageClass.rawValue {
            keepEvent()
        } else if CheckStatus.NoKeep.rawValue == viewPageClass.rawValue {
            noKeepEvent()
        }
        
        // イベントの情報
        if self.eventSummary.checkStatus == CheckStatus.NoCheck.rawValue {
            noCheckEvent()
        } else if self.eventSummary.checkStatus == CheckStatus.Keep.rawValue {
            keepEvent()
        } else if self.eventSummary.checkStatus == CheckStatus.NoKeep.rawValue {
            noKeepEvent()
        }
        
        if eventSummary.apiId == ApiId.Atdn.rawValue {
            apiNameLbl.text = apiFromMessage + ApiId.Atdn.getName()
        } else if eventSummary.apiId == ApiId.Connpass.rawValue {
            apiNameLbl.text = apiFromMessage + ApiId.Connpass.getName()
        } else if eventSummary.apiId == ApiId.Doorkeeper.rawValue {
            apiNameLbl.text = apiFromMessage + ApiId.Doorkeeper.getName()
        }
        
        eventStatusLbl.text = "\(eventSummary.accepted) / 定員\(eventSummary.limit)人"
        
        self.indexPath = indexPath
        
        titileLbl.text = eventSummary.title
        titileLbl.numberOfLines = 0
        titileLbl.lineBreakMode = .ByWordWrapping
        
        addressLbl.text = eventSummary.address != "" ? eventSummary.address : "開催地未定"
        eventSummary.eventDate = ViewFormaatter.sharedInstance.setEventDate(eventSummary)
        dateLbl.text = eventSummary.eventDate
        
        id = eventSummary.id
    }
    
}
