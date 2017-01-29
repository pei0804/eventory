//
//  CheckListTableViewCell.swift
//  Eventory
//
//  Created by jumpei on 2016/09/06.
//  Copyright © 2016年 jumpei. All rights reserved.
//

import UIKit

class CheckListTableViewCell: UITableViewCell {

    override func awakeFromNib() {
        super.awakeFromNib()
        self.tintColor = Colors.main
        self.separatorInset = UIEdgeInsetsZero
        self.layoutMargins = UIEdgeInsetsZero
    }

    override func setSelected(selected: Bool, animated: Bool) {
        super.setSelected(selected, animated: animated)
    }
    
    func bind(userRegister: Dictionary<String, AnyObject>) {
        self.textLabel?.text = userRegister["name"] as? String
        if userRegister["status"] as! Bool {
            self.check()
        } else {
            self.checkRemove()
        }
    }
    
    func checkAction(inout userRegister: [Dictionary<String, AnyObject>]?, indexPath: NSIndexPath, inout checkCount: Int) {
        if self.accessoryType == .None {
            userRegister![indexPath.row]["status"] = true
            checkCount += 1
            self.check()
        } else {
            userRegister![indexPath.row]["status"] = false
            checkCount -= 1
            self.checkRemove()
        }
    }
    
    func check() {
        self.accessoryType = .Checkmark
        self.textLabel?.font = UIFont.boldSystemFontOfSize(17)
        self.textLabel?.textColor = Colors.main
    }
    
    func checkRemove() {
        self.accessoryType = .None
        self.textLabel?.font = UIFont.systemFontOfSize(17)
        self.textLabel?.textColor = UIColor.blackColor()
    }
}
