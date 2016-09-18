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
    }

    override func setSelected(selected: Bool, animated: Bool) {
        
        super.setSelected(selected, animated: animated)
    }
    
    func bind(userRegister: Dictionary<String, AnyObject>) {
        
        self.textLabel?.text = userRegister["name"] as? String
        if userRegister["status"] as! Bool {
            check()
        } else {
            checkRemove()
        }
    }
    
    func checkAction(inout userRegister: [Dictionary<String, AnyObject>]?, indexPath: NSIndexPath, inout checkCount: Int) {
        
        if self.accessoryType == .None {
            userRegister![indexPath.row]["status"] = true
            checkCount += 1
            check()
        } else {
            userRegister![indexPath.row]["status"] = false
            checkCount -= 1
            checkRemove()
        }
    }
    
    func check() {
        
        self.accessoryType = .Checkmark
        self.textLabel?.font = UIFont.boldSystemFontOfSize(17)
        self.textLabel?.textColor = UIColor.blueColor()
    }
    
    func checkRemove() {
        
        self.accessoryType = .None
        self.textLabel?.font = UIFont.systemFontOfSize(17)
        self.textLabel?.textColor = UIColor.blackColor()
    }
}
