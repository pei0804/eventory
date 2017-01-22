//
//  BaseTableViewController.swift
//  Eventory
//
//  Created by jumpei on 2017/01/22.
//  Copyright © 2017年 jumpei. All rights reserved.
//

import UIKit
import DZNEmptyDataSet
import SafariServices
import SVProgressHUD

class BaseTableViewController: UITableViewController, SFSafariViewControllerDelegate, DZNEmptyDataSetSource, DZNEmptyDataSetDelegate {

    @IBOutlet weak var scrollView: UIScrollView?

    var eventSummaries: [EventSummary]? {
        didSet {
            if let eventSummaries = self.eventSummaries where eventSummaries.count == 0 {
                self.tableView.setContentOffset(CGPointZero, animated: false)
            }
            self.tableView.reloadData()
        }
    }

    override func viewDidLoad() {
        super.viewDidLoad()

        self.scrollView = tableView
        self.addRefreshControl()
        self.refresh()

        self.tableView.emptyDataSetSource = self
        self.tableView.emptyDataSetDelegate = self

        self.edgesForExtendedLayout = UIRectEdge.None

        self.tableView.registerNib(UINib(nibName: EventInfoTableViewCellIdentifier, bundle: nil), forCellReuseIdentifier: EventInfoTableViewCellIdentifier)
        NSNotificationCenter.defaultCenter().addObserver(self, selector: #selector(self.becomeActive(_:)), name: UIApplicationDidBecomeActiveNotification, object: nil)
    }

    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
        // Dispose of any resources that can be recreated.
    }

    deinit {
        self.tableView.delegate = nil
        NSNotificationCenter.defaultCenter().removeObserver(self)
    }

    var refreshControlY: CGFloat = 3.0

    func addRefreshControl() {
        if let scrollView = self.scrollView {
            let refreshControl = UIRefreshControl()
            refreshControl.attributedTitle = NSAttributedString(string: "更新")
            refreshControl.addTarget(self, action: #selector(BaseViewController.pullRefresh(_:)), forControlEvents: .ValueChanged)
            if let tableView = scrollView as? UITableView {
                tableView.backgroundView = refreshControl
                tableView.alwaysBounceVertical = true
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
        SVProgressHUD.showWithStatus(ServerConnectionMessage)
        self.refresh() {
            SVProgressHUD.dismiss()
        }
    }

    @IBAction func pullRefresh(refreshControl: UIRefreshControl) {

        SVProgressHUD.showWithStatus(ServerConnectionMessage)
        self.handleRefresh()
        self.refresh() {
            SVProgressHUD.dismiss()
            refreshControl.endRefreshing()
        }
    }

    func refresh(completed: (() -> Void)? = nil) {
        dispatch_async(dispatch_get_main_queue()) {
            completed?()
        }
    }

    // MARK: - UITableViewDataSource

    override func numberOfSectionsInTableView(tableView: UITableView) -> Int {
        return 1
    }

    override func tableView(tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        if let eventSummaries = self.eventSummaries {
            return eventSummaries.count
        }
        return 0
    }

    override func tableView(tableView: UITableView, heightForRowAtIndexPath indexPath: NSIndexPath) -> CGFloat {
        return EventInfoCellHeight
    }


    // MARK: - UITableViewDelegate
    
    override func tableView(tableView: UITableView, didSelectRowAtIndexPath indexPath:NSIndexPath) {

        guard let eventSummaries = self.eventSummaries else {
            return
        }
        let url: String = eventSummaries[indexPath.row].url
        if !url.lowercaseString.hasPrefix("http://") && !url.lowercaseString.hasPrefix("https://") {
            let alert: UIAlertController = UIAlertController(title: "不正なリンクを検出しました", message: "このイベントに設定されているリンクに問題がありました。", preferredStyle: .Alert)
            let cancelAction: UIAlertAction = UIAlertAction(title: "OK", style: .Cancel, handler: nil)
            alert.addAction(cancelAction)
            self.presentViewController(alert, animated: true, completion: nil)
            tableView.deselectRowAtIndexPath(indexPath, animated: true)
            return
        }
        let brow = SFSafariViewController(URL: NSURL(string: url)!, entersReaderIfAvailable: false)
        brow.delegate = self
        presentViewController(brow, animated: true, completion: nil)
    }

    // MARK: - DZNEmptyDataSetSource

    func titleForEmptyDataSet(scrollView: UIScrollView!) -> NSAttributedString! {
        let text = "条件に合致する情報がありません"
        let attribs = [
            NSFontAttributeName: UIFont.boldSystemFontOfSize(18),
            NSForegroundColorAttributeName: UIColor.darkGrayColor()
        ]

        return NSAttributedString(string: text, attributes: attribs)
    }

    // MARK: - DZNEmptyDataSetDelegate

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
