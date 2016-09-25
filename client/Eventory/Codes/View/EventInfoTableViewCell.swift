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
    }
    
    override func setSelected(selected: Bool, animated: Bool) {
        super.setSelected(selected, animated: animated)
    }
    
    @IBOutlet weak var noKeepBtn: UIButton!
    @IBOutlet weak var keepBtn: UIButton!
    @IBOutlet weak var dateLbl: UILabel!
    
    @IBOutlet weak var addressLbl: UILabel!
    @IBOutlet weak var titileLbl: UILabel!
    @IBOutlet weak var titleBackgroundView: UIView!
    
    var id: Int = 0
    var indexPath: NSIndexPath = NSIndexPath(index: 0)
    var eventSummary: EventSummary = EventSummary()
    
    @IBAction func keepBtn(sender: AnyObject) {
        EventManager.sharedInstance.keepAction(id, isKeep: true)
        keepBtn.setTitle("キープ◯", forState: .Normal)
        noKeepBtn.setTitle("興味なし", forState: .Normal)
        titleBackgroundView.backgroundColor = Colors.main
        eventSummary.checkStatus = CheckStatus.Keep.rawValue
    }
    
    @IBAction func noKeepBtn(sender: AnyObject) {
        EventManager.sharedInstance.keepAction(id, isKeep: false)
        noKeepBtn.setTitle("興味なし×", forState: .Normal)
        keepBtn.setTitle("キープ", forState: .Normal)
        titleBackgroundView.backgroundColor = Colors.noKeep
        eventSummary.checkStatus = CheckStatus.NoKeep.rawValue
    }
    
    func bind(eventSummary: EventSummary, viewPageClass: CheckStatus, indexPath: NSIndexPath) {
        
        self.eventSummary = eventSummary
        
        id = eventSummary.id
        if eventSummary.checkStatus == CheckStatus.NoCheck.rawValue {
            keepBtn.setTitle("キープ", forState: .Normal)
            noKeepBtn.setTitle("興味なし", forState: .Normal)
            titleBackgroundView.backgroundColor = Colors.noCheck
        }
        else if eventSummary.checkStatus == CheckStatus.Keep.rawValue {
            keepBtn.setTitle("キープ◯", forState: .Normal)
            noKeepBtn.setTitle("興味なし", forState: .Normal)
            titleBackgroundView.backgroundColor = Colors.main
        }
        else if eventSummary.checkStatus == CheckStatus.NoKeep.rawValue {
            noKeepBtn.setTitle("興味なし×", forState: .Normal)
            keepBtn.setTitle("キープ", forState: .Normal)
            titleBackgroundView.backgroundColor = Colors.noKeep
        }
        
        if CheckStatus.Keep.rawValue == viewPageClass.rawValue {
            keepBtn.hidden = true
            titleBackgroundView.backgroundColor = Colors.main
        }
        else if CheckStatus.NoKeep.rawValue == viewPageClass.rawValue {
            noKeepBtn.hidden = true
            titleBackgroundView.backgroundColor = Colors.noKeep
        }
        
        self.indexPath = indexPath
        
        titileLbl.text = eventSummary.title
        //descLbl.text = eventSummary.desc
        addressLbl.text = eventSummary.address != "" ? eventSummary.address : "未定"
        eventSummary.eventDate = ViewFormaatter.sharedInstance.setEventDate(eventSummary)
        dateLbl.text = eventSummary.eventDate
        
        noKeepBtn.tag = eventSummary.id
        keepBtn.tag = eventSummary.id
    }
    
}
