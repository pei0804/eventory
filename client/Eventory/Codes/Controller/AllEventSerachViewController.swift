//
//  AllEventSerachViewController.swift
//  Eventory
//
//  Created by jumpei on 2016/09/18.
//  Copyright © 2016年 jumpei. All rights reserved.
//

import UIKit
import SafariServices
import SwiftTask
import SVProgressHUD

class AllEventSerachViewController: BaseTableViewController {

    private var freeWordSearchBar: BaseSearchBar!

    override var viewPageClass: CheckStatus {
        return CheckStatus.Search
    }

    override func viewDidLoad() {
        super.viewDidLoad()

        // 検索バーを作成する.
        freeWordSearchBar = BaseSearchBar()
        freeWordSearchBar.awakeFromNib()
        self.freeWordSearchBar.delegate = self
        self.freeWordSearchBar.layer.position = CGPoint(x: self.view.bounds.width/2, y: 100)
        self.freeWordSearchBar.placeholder = "検索 （例：Swift 東京）"

        self.view.addSubview(freeWordSearchBar)
        self.tableView.registerNib(UINib(nibName: EventInfoTableViewCellIdentifier, bundle: nil), forCellReuseIdentifier: EventInfoTableViewCellIdentifier)
    }
    
    override func tableView(tableView: UITableView, viewForHeaderInSection section: Int) -> UIView? {
        return freeWordSearchBar
    }

    override func tableView(tableView: UITableView, heightForHeaderInSection section: Int) -> CGFloat {
        return 40.0;
    }
}

// MARK: - UISearchBarDelegate 

extension AllEventSerachViewController: UISearchBarDelegate {
    
    func searchBarCancelButtonClicked(searchBar: UISearchBar) {
        self.freeWordSearchBar.text = ""
        self.freeWordSearchBar.resignFirstResponder()
    }
    
    func searchBarSearchButtonClicked(searchBar: UISearchBar) {
        let term = self.freeWordSearchBar.text ?? ""
        if !term.isEmpty {
            self.eventSummaries = EventManager.sharedInstance.getNewEventAll(term)
        }
        self.freeWordSearchBar.resignFirstResponder()
    }
    
    override func touchesBegan(touches: Set<UITouch>, withEvent event: UIEvent?) {
        if (self.freeWordSearchBar.isFirstResponder()) {
            self.freeWordSearchBar.resignFirstResponder()
        }
    }
}
