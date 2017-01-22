//
//  KeepButton.swift
//  Eventory
//
//  Created by jumpei on 2016/12/30.
//  Copyright © 2016年 jumpei. All rights reserved.
//

import UIKit

class KeepButton: UIButton {

    override func awakeFromNib() {
        super.awakeFromNib()
        self.setup()
    }
    
    private func setup() {
        self.layer.cornerRadius = 4.0
    }
    
    func active() {
        self.layer.backgroundColor = Colors.main.CGColor
        self.layer.borderColor = UIColor.clearColor().CGColor;
        self.layer.borderWidth = 0;
        self.setTitleColor(UIColor.whiteColor(), forState: .Normal)
        self.setImage(UIImage(named:"keepActive.png"), forState: .Normal)
        
    }
    
    func noActive() {
        self.layer.backgroundColor = UIColor.clearColor().CGColor
        self.layer.borderColor = Colors.main.CGColor;
        self.layer.borderWidth = 2;
        self.setTitleColor(Colors.main, forState: .Normal)
        self.setImage(UIImage(named:"keepNoActive.png"), forState: .Normal)
    }
}
