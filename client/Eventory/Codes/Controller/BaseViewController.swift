//
//  BaseViewController.swift
//  Eventory
//
//  Created by jumpei on 2016/09/18.
//  Copyright © 2016年 jumpei. All rights reserved.
//

import UIKit
import DZNEmptyDataSet

class BaseViewController: UIViewController {
    
    weak var refreshControl: UIRefreshControl?
    @IBOutlet weak var scrollView: UIScrollView?
    
    override func viewDidLoad() {
        
        super.viewDidLoad()
        NSNotificationCenter.defaultCenter().addObserver(self, selector: #selector(self.becomeActive(_:)), name: UIApplicationDidBecomeActiveNotification, object: nil)
    }
    
    override func didReceiveMemoryWarning() {
        
        super.didReceiveMemoryWarning()
    }
    
    var refreshControlY: CGFloat = 3.0
    
    func addRefreshControl() {
        
        if let scrollView = self.scrollView {
            let refreshControl = UIRefreshControl()
            refreshControl.attributedTitle = NSAttributedString(string: "更新")
            refreshControl.addTarget(self, action: #selector(BaseViewController.pullRefresh(_:)), forControlEvents: .ValueChanged)
            if let tableView = scrollView as? UITableView {
                tableView.backgroundView = refreshControl
            }
            refreshControl.bounds.origin.y = -self.refreshControlY
            scrollView.alwaysBounceVertical = true
            scrollView.setContentOffset(CGPoint(x: 0, y: -refreshControl.frame.size.height-10), animated: true)
            self.refreshControl = refreshControl
        }
    }
    
    
    func handleRefresh() {
    }

    func becomeActive(notification: NSNotification) {
    }

    @IBAction func pullRefresh(refreshControl: UIRefreshControl) {
        
        self.handleRefresh()
        self.refresh() {
            refreshControl.endRefreshing()
        }
    }
    
    func refresh(completed: (() -> Void)? = nil) {
        
        dispatch_async(dispatch_get_main_queue()) {
            completed?()
        }
    }

    deinit {
        NSNotificationCenter.defaultCenter().removeObserver(self)
    }
}

// MARK: - DZNEmptyDataSetSource

extension BaseViewController: DZNEmptyDataSetSource {
    
    func titleForEmptyDataSet(scrollView: UIScrollView!) -> NSAttributedString! {
        
        let text = "条件に合致する情報がありません"
        let attribs = [
            NSFontAttributeName: UIFont.boldSystemFontOfSize(18),
            NSForegroundColorAttributeName: UIColor.darkGrayColor()
        ]
        
        return NSAttributedString(string: text, attributes: attribs)
    }
}

// MARK: - DZNEmptyDataSetDelegate

extension BaseViewController: DZNEmptyDataSetDelegate {
    
    func emptyDataSetShouldAllowScroll(scrollView: UIScrollView!) -> Bool {
        
        return true
    }
    
    func emptyDataSetWillAppear(scrollView: UIScrollView!) {
        
        if let tableView = self.scrollView as? UITableView {
            tableView.separatorColor = UIColor.clearColor();
        }
    }
    
    func emptyDataSetDidDisappear(scrollView: UIScrollView!) {
        
        if let tableView = self.scrollView as? UITableView {
            tableView.separatorColor = UIColor.grayColor();
        }
    }
}
