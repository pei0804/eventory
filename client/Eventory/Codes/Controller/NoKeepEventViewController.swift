//
//  NoKeepEventViewController.swift
//  Eventory
//
//  Created by jumpei on 2016/08/22.
//  Copyright © 2016年 jumpei. All rights reserved.
//

import UIKit
import SVProgressHUD
import SwiftTask

class NoKeepEventViewController: BaseTableViewController {

    override var viewPageClass: CheckStatus {
        return CheckStatus.NoKeep
    }

    override func viewDidLoad() {
        super.viewDidLoad()
    }

    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
    }
}
