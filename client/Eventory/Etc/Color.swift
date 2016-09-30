//
//  Color.swift
//  Eventory
//
//  Created by jumpei on 2016/09/19.
//  Copyright © 2016年 jumpei. All rights reserved.
//

import UIKit

struct Colors {
    
    static let main       = UIColorFromRGB(0xE67E22)
    static let main_bg    = UIColorFromRGB(0xFFE8D3)
    
    
    static let noKeep     = UIColorFromRGB(0x95a5a6)
    static let noKeep_bg  = UIColorFromRGB(0xEAEAEA)
    static let noKeep_btn = UIColorFromRGB(0xD8D8D8)
    
    static let noCheck    = UIColorFromRGB(0x3498db)
    static let noCheck_bg = UIColorFromRGB(0xF2F7FF)
    
    static let glay     = UIColorFromRGB(0x9B9B9B)
}


func UIColorFromRGB(rgbValue: UInt) -> UIColor {
    return UIColor(
        red:    CGFloat((rgbValue & 0xFF0000) >> 16) / 255.0,
        green:  CGFloat((rgbValue & 0x00FF00) >> 8) / 255.0,
        blue:   CGFloat(rgbValue & 0x0000FF) / 255.0,
        alpha:  CGFloat(1.0)
    )
}
