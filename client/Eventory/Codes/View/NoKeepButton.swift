//
//  NoKeepButton.swift
//  Eventory
//
//  Created by jumpei on 2016/12/31.
//  Copyright © 2016年 jumpei. All rights reserved.
//

import UIKit

class NoKeepButton: UIButton {

    override func awakeFromNib() {
        super.awakeFromNib()
        self.layer.cornerRadius = 4.0
    }
    
    func active() {
        self.layer.borderColor = UIColor.clearColor().CGColor;
        self.layer.backgroundColor = Colors.noKeep.CGColor
        self.layer.borderWidth = 0;
        self.setTitleColor(UIColor.whiteColor(), forState: .Normal)
        self.setImage(UIImage(named:"noKeepActive.png"), forState: .Normal)
    }
    
    func noActive() {
        self.layer.backgroundColor = UIColor.clearColor().CGColor
        self.layer.borderColor = Colors.noKeep.CGColor;
        self.layer.borderWidth = 2;
        self.setTitleColor(Colors.noKeep, forState: .Normal)
        self.setImage(UIImage(named:"noKeepNoActive.png"), forState: .Normal)
    }
}
