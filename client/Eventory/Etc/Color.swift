//
//  Color.swift
//  Eventory
//
//  Created by jumpei on 2016/09/19.
//  Copyright © 2016年 jumpei. All rights reserved.
//

import UIKit

struct Colors {
    
    // 現状は色にネーミングで意味を持たせない
    // いずれも番号がついていないものが一番濃い色にしている
    static let main     = UIColorFromRGB(0xE67E22)
    static let main2    = UIColorFromRGB(0xFFE8D3)
    
    static let noKeep   = UIColorFromRGB(0x95a5a6)
    static let noKeep2  = UIColorFromRGB(0xEAEAEA)
    static let noKeep3  = UIColorFromRGB(0xD8D8D8)
    
    static let noCheck  = UIColorFromRGB(0x3498db)
    static let noCheck2 = UIColorFromRGB(0xF2F7FF)
}


func UIColorFromRGB(rgbValue: UInt) -> UIColor {
    
    return UIColor(
        red:    CGFloat((rgbValue & 0xFF0000) >> 16) / 255.0,
        green:  CGFloat((rgbValue & 0x00FF00) >> 8) / 255.0,
        blue:   CGFloat(rgbValue & 0x0000FF) / 255.0,
        alpha:  CGFloat(1.0)
    )
}
