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
        descLbl.sizeToFit()
    }
    
    override func setSelected(selected: Bool, animated: Bool) {
        super.setSelected(selected, animated: animated)
    }
    
    @IBOutlet weak var noKeepBtn: UIButton!
    @IBOutlet weak var keepBtn: UIButton!
    @IBOutlet weak var dateLbl: UILabel!
    
    @IBOutlet weak var addressLbl: UILabel!
    @IBOutlet weak var descLbl: UILabel!
    @IBOutlet weak var titileLbl: UILabel!
    
    var id: Int = 0
    var indexPath: NSIndexPath = NSIndexPath(index: 0)
    var eventSummary: EventSummary = EventSummary()
    
    @IBAction func keepBtn(sender: AnyObject) {
        EventManager.sharedInstance.keepAction(id, isKeep: true)
        keepBtn.setTitle("キープ◯", forState: .Normal)
        noKeepBtn.setTitle("興味なし", forState: .Normal)
        self.eventSummary.checkStatus = CheckStatus.Keep.rawValue
    }
    
    @IBAction func noKeepBtn(sender: AnyObject) {
        EventManager.sharedInstance.keepAction(id, isKeep: false)
        noKeepBtn.setTitle("興味なし×", forState: .Normal)
        keepBtn.setTitle("キープ", forState: .Normal)
        self.eventSummary.checkStatus = CheckStatus.NoKeep.rawValue
    }
    
    func bind(eventSummary: EventSummary, viewPageClass: CheckStatus, indexPath: NSIndexPath) {
        
        self.eventSummary = eventSummary
        
        id = eventSummary.id
        if eventSummary.checkStatus == CheckStatus.NoCheck.rawValue {
            self.keepBtn.setTitle("キープ", forState: .Normal)
            self.noKeepBtn.setTitle("興味なし", forState: .Normal)
        }
        else if eventSummary.checkStatus == CheckStatus.Keep.rawValue {
            self.keepBtn.setTitle("キープ◯", forState: .Normal)
            self.noKeepBtn.setTitle("興味なし", forState: .Normal)
        }
        else if eventSummary.checkStatus == CheckStatus.NoKeep.rawValue {
            self.noKeepBtn.setTitle("興味なし×", forState: .Normal)
            self.keepBtn.setTitle("キープ", forState: .Normal)
        }
        
        if CheckStatus.Keep.rawValue == viewPageClass.rawValue {
            keepBtn.hidden = true
        }
        else if CheckStatus.NoKeep.rawValue == viewPageClass.rawValue {
            noKeepBtn.hidden = true
        }
        
        self.indexPath = indexPath
        
        titileLbl.text = eventSummary.title
        descLbl.text = eventSummary.desc
        addressLbl.text = eventSummary.address != "" ? eventSummary.address : "未定"
        eventSummary.eventDate = ViewFormaatter.sharedInstance.setEventDate(eventSummary)
        dateLbl.text = eventSummary.eventDate
        
        noKeepBtn.tag = eventSummary.id
        keepBtn.tag = eventSummary.id
    }
    
}
