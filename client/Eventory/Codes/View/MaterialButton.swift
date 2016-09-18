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
        layer.masksToBounds = true
        // 円を描画
        let shapeLayer = CAShapeLayer()
        shapeLayer.fillColor = UIColor.grayColor().CGColor
        shapeLayer.path = UIBezierPath(ovalInRect: tapEffectView.bounds).CGPath
        tapEffectView.layer.addSublayer(shapeLayer)
        tapEffectView.hidden = true
        
        addSubview(tapEffectView)
    }
    
    override func touchesBegan(touches: Set<UITouch>, withEvent event: UIEvent?) {
        super.touchesBegan(touches, withEvent: event)
        if let point = touches.first?.locationInView(self) {
            tapEffectView.frame.origin = point
            
            tapEffectView.alpha = 1.0
            tapEffectView.hidden = false
            tapEffectView.transform = CGAffineTransformMakeScale(1.0, 1.0)
            
            UIView.animateWithDuration(0.5,
                                       animations: {
                                        self.tapEffectView.alpha = 0
                                        self.tapEffectView.transform = CGAffineTransformMakeScale(200.0, 200.0)
            }) { finished in
                self.tapEffectView.hidden = true
                self.tapEffectView.transform = CGAffineTransformMakeScale(1.0, 1.0)
            }
        }
    }
}




