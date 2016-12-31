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
        setup()
    }
    
    private func setup() {
        layer.cornerRadius = 4.0
    }
    
    func active() {
        layer.backgroundColor = Colors.main.CGColor
        layer.borderColor = UIColor.clearColor().CGColor;
        layer.borderWidth = 0;
        setTitleColor(UIColor.whiteColor(), forState: .Normal)
    }
    
    func noActive() {
        layer.backgroundColor = UIColor.clearColor().CGColor
        layer.borderColor = Colors.main.CGColor;
        layer.borderWidth = 2;
        setTitleColor(Colors.main, forState: .Normal)
    }
    
    enum Status: Int {
        case noActive = 0
        case Active = 1
        case None = 5
    }
}
