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
        setup()
    }
    
    private func setup() {
        layer.cornerRadius = 4.0
    }
    
    func active() {
        layer.backgroundColor = Colors.noKeep.CGColor
        layer.borderColor = UIColor.clearColor().CGColor;
        layer.borderWidth = 0;
        setTitleColor(UIColor.whiteColor(), forState: .Normal)
    }
    
    func noActive() {
        layer.backgroundColor = UIColor.clearColor().CGColor
        layer.borderColor = Colors.noKeep.CGColor;
        layer.borderWidth = 2;
        setTitleColor(Colors.noKeep, forState: .Normal)
    }
    
    enum Status: Int {
        case noActive = 0
        case Active = 1
        case None = 5
    }
}
