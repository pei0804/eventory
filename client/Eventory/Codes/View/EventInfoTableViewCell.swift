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
    
    @IBOutlet weak var noKeepButton: NoKeepButton!
    @IBOutlet weak var keepButton: KeepButton!
    @IBOutlet weak var dateLbl: UILabel!
    @IBOutlet weak var apiNameLbl: UILabel!
    @IBOutlet weak var eventStatusLbl: UILabel!
    @IBOutlet weak var addressLbl: UILabel!
    @IBOutlet weak var titileLbl: UILabel!
    @IBOutlet weak var titleLblWidth: NSLayoutConstraint!
    @IBOutlet weak var titleBackgroundView: UIView!
    
    var id: Int = 0
    var indexPath: NSIndexPath = NSIndexPath(index: 0)
    var eventSummary: EventSummary = EventSummary()
    let apiFromMessage = "情報提供元："
    
    @IBAction func keepBtn(sender: AnyObject) {
        EventManager.sharedInstance.keepAction(id, isKeep: true)
        eventSummary.checkStatus = CheckStatus.Keep.rawValue
        titileLbl.textColor = Colors.main
        titleBackgroundView.backgroundColor = Colors.main_bg
    }
    
    @IBAction func noKeepBtn(sender: AnyObject) {
        EventManager.sharedInstance.keepAction(id, isKeep: false)
        eventSummary.checkStatus = CheckStatus.NoKeep.rawValue
        titileLbl.textColor = Colors.noKeep
        titleBackgroundView.backgroundColor = Colors.noKeep_bg
    }
    
    func bind(eventSummary: EventSummary, viewPageClass: CheckStatus, indexPath: NSIndexPath) {
        
        if eventSummary.checkStatus == CheckStatus.NoCheck.rawValue {
            titileLbl.textColor = Colors.noCheck
            titleBackgroundView.backgroundColor = Colors.noCheck_bg
        } else if eventSummary.checkStatus == CheckStatus.Keep.rawValue {
            titileLbl.textColor = Colors.main
            titleBackgroundView.backgroundColor = Colors.main_bg
        } else if eventSummary.checkStatus == CheckStatus.NoKeep.rawValue {
            titileLbl.textColor = Colors.noKeep
            titleBackgroundView.backgroundColor = Colors.noKeep_bg
        }
        
        if CheckStatus.Keep.rawValue == viewPageClass.rawValue {
            keepBtn.hidden = true
            titileLbl.textColor = Colors.main
            titleBackgroundView.backgroundColor = Colors.main_bg
        } else if CheckStatus.NoKeep.rawValue == viewPageClass.rawValue {
            noKeepBtn.hidden = true
            titileLbl.textColor = Colors.noKeep
            titleBackgroundView.backgroundColor = Colors.noKeep_bg
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
        let frame = CGSizeMake(300, CGFloat.max)
        let rect = titileLbl.sizeThatFits(frame)
        titleLblWidth.constant = rect.width
        
        addressLbl.text = eventSummary.address != "" ? eventSummary.address : "開催地未定"
        eventSummary.eventDate = ViewFormaatter.sharedInstance.setEventDate(eventSummary)
        dateLbl.text = eventSummary.eventDate
        
        id = eventSummary.id
    }
    
}
