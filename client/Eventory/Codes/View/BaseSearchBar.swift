//
//  BaseSearchBar.swift
//  Eventory
//
//  Created by jumpei on 2017/01/24.
//  Copyright © 2017年 jumpei. All rights reserved.
//

import UIKit

class BaseSearchBar: UISearchBar {

    override func awakeFromNib() {
        super.awakeFromNib()
        self.layer.cornerRadius = 4.0
        self.frame = CGRectMake(0, 0, 320, 40)
        self.tintColor = Colors.main
    }
}



