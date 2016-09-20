//
//  EventInfoViewController.swift
//  Eventory
//
//  Created by jumpei on 2016/08/19.
//  Copyright © 2016年 jumpei. All rights reserved.
//

import UIKit
import SafariServices
import SwiftTask
import DZNEmptyDataSet

class EventInfoViewController: BaseViewController, SFSafariViewControllerDelegate {
    
    var eventSummarys: [EventSummary]? {
        didSet {
            self.tableView.reloadData()
        }
    }
    
    @IBOutlet weak var tableView: UITableView!
    
    override func viewDidLoad() {
        super.viewDidLoad()
        
        self.scrollView = tableView
        self.addRefreshControl()
        
        tableView.delegate = self
        tableView.dataSource = self
        
        self.tableView.emptyDataSetSource = self
        
        self.tableView.registerNib(UINib(nibName: EventInfoTableViewCellIdentifier, bundle: nil), forCellReuseIdentifier: EventInfoTableViewCellIdentifier)
        // TODO: パス確認用（削除必須）
        // print(NSSearchPathForDirectoriesInDomains(.DocumentDirectory, .UserDomainMask, true))
    }
    
    override func viewWillAppear(animated:Bool) {
        
        super.viewWillAppear(animated)
        eventSummarys = EventManager.sharedInstance.getSelectNewEventAll()
        let appDelegate:AppDelegate = UIApplication.sharedApplication().delegate as! AppDelegate
        appDelegate.newEvent = EventManager.sharedInstance.getSelectNewEventAll().count
    }
    
    override func viewWillDisappear(animated:Bool) {
        
        super.viewWillDisappear(animated)
        let appDelegate:AppDelegate = UIApplication.sharedApplication().delegate as! AppDelegate
        appDelegate.newEvent = EventManager.sharedInstance.getSelectNewEventAll().count
    }
    
    override func didReceiveMemoryWarning() {
        
        super.didReceiveMemoryWarning()
    }
    
    override func refresh(completed: (() -> Void)? = nil) {
        dispatch_async(dispatch_get_main_queue()) {
            EventManager.sharedInstance.getNewEventAll()
            self.eventSummarys = EventManager.sharedInstance.getSelectNewEventAll()
            completed?()
        }
    }
}



// MARK: - UITableViewDataSource

extension EventInfoViewController: UITableViewDataSource {
    
    func numberOfSectionsInTableView(tableView: UITableView) -> Int {
        
        return 1
    }
    
    func tableView(tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        
        if let eventSummarys = eventSummarys {
            return eventSummarys.count
        }
        return 0
    }
    
    func tableView(tableView: UITableView, cellForRowAtIndexPath indexPath: NSIndexPath) -> UITableViewCell {
        
        if let cell = tableView.dequeueReusableCellWithIdentifier(EventInfoTableViewCellIdentifier, forIndexPath: indexPath) as? EventInfoTableViewCell {
            if let eventSummarys = eventSummarys {
                cell.bind(eventSummarys[indexPath.row], viewPageClass: CheckStatus.NoCheck, indexPath: indexPath)
                return cell
            }
        }
        return UITableViewCell()
    }
    
    func tableView(tableView: UITableView, heightForRowAtIndexPath indexPath: NSIndexPath) -> CGFloat {
        return 250
    }
}

// MARK: - UITableViewDelegate

extension EventInfoViewController: UITableViewDelegate {
    
    func tableView(tableView: UITableView, didSelectRowAtIndexPath indexPath:NSIndexPath) {
        
        guard let eventSummarys = eventSummarys else {
            return
        }
        let url: String = eventSummarys[indexPath.row].url
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
}


extension EventInfoViewController: DZNEmptyDataSetSource {
    
    func titleForEmptyDataSet(scrollView: UIScrollView!) -> NSAttributedString! {
        let text = "条件に合致する情報がありません"
        let attribs = [
            NSFontAttributeName: UIFont.boldSystemFontOfSize(18),
            NSForegroundColorAttributeName: UIColor.darkGrayColor()
        ]
        
        return NSAttributedString(string: text, attributes: attribs)
    }
}
