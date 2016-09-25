//
//  MaterialButton.swift
//  Eventory
//
//  Created by jumpei on 2016/08/22.
//  Copyright © 2016年 jumpei. All rights reserved.
//

import UIKit

class MaterialButton: UIButton {
    
    private let tapEffectView = UIView(frame: CGRect(x: 0, y: 0, width: 1, height: 1))
    
    override func awakeFromNib() {
        super.awakeFromNib()
        
        setup()
    }
    
    private func setup() {
        
        // ボタン自体を角丸にする
        layer.cornerRadius = 4.0
    }
}




