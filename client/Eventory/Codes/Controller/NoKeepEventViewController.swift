//
//  NoKeepEventViewController.swift
//  Eventory
//
//  Created by jumpei on 2016/08/22.
//  Copyright © 2016年 jumpei. All rights reserved.
//

import UIKit
import SafariServices
import SwiftTask
import SVProgressHUD

class NoKeepEventViewController: BaseViewController {

    var eventSummaries: [EventSummary]? {
        didSet {
            if let eventSummaries = eventSummaries where eventSummaries.count == 0 {
                tableView.setContentOffset(CGPointZero, animated: false)
            }
            self.tableView.reloadData()
        }
    }
    
    @IBOutlet weak var tableView: UITableView!
    
    override func viewDidLoad() {
        super.viewDidLoad()
        
        scrollView = tableView
        addRefreshControl()
        
        tableView.emptyDataSetSource = self;
        tableView.emptyDataSetDelegate = self;
        
        tableView.delegate = self
        tableView.dataSource = self
        
        self.tableView.registerNib(UINib(nibName: EventInfoTableViewCellIdentifier, bundle: nil), forCellReuseIdentifier: EventInfoTableViewCellIdentifier)
    }
    
    override func viewWillAppear(animated:Bool) {
        
        super.viewWillAppear(animated)
        eventSummaries = EventManager.sharedInstance.getNoKeepEventAll()
    }
    
    override func didReceiveMemoryWarning() {
        
        super.didReceiveMemoryWarning()
    }
    
    override func refresh(completed: (() -> Void)? = nil) {
        dispatch_async(dispatch_get_main_queue()) {

            SVProgressHUD.showWithStatus(ServerConnectionMessage)
            let task = [EventManager.sharedInstance.fetchNewEvent()]
            
            Task.all(task).success { _ in
                SVProgressHUD.dismiss()
                self.eventSummaries = EventManager.sharedInstance.getNoKeepEventAll()
                completed?()
                }.failure { _ in
                    SVProgressHUD.dismiss()
                    let alert: UIAlertController = UIAlertController(title: NetworkErrorTitle,message: NetworkErrorMessage, preferredStyle: .Alert)
                    let cancelAction: UIAlertAction = UIAlertAction(title: NetworkErrorButton, style: .Cancel, handler: nil)
                    alert.addAction(cancelAction)
                    self.presentViewController(alert, animated: true, completion: nil)
                    completed?()
            }
        }
    }
}


// MARK: - UITableViewDataSource

extension NoKeepEventViewController: UITableViewDataSource {
    
    func numberOfSectionsInTableView(tableView: UITableView) -> Int {
        
        return 1
    }
    
    func tableView(tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        
        if let eventSummaries = eventSummaries {
            return eventSummaries.count
        }
        return 0
    }
    
    func tableView(tableView: UITableView, cellForRowAtIndexPath indexPath: NSIndexPath) -> UITableViewCell {
        
        if let cell = tableView.dequeueReusableCellWithIdentifier(EventInfoTableViewCellIdentifier, forIndexPath: indexPath) as? EventInfoTableViewCell {
            if let eventSummaries = eventSummaries {
                cell.bind(eventSummaries[indexPath.row], viewPageClass: CheckStatus.NoKeep,indexPath: indexPath)
                return cell
            }
        }
        return UITableViewCell()
    }
    
    func tableView(tableView: UITableView, heightForRowAtIndexPath indexPath: NSIndexPath) -> CGFloat {
        return EventInfoCellHeight
    }
}

// MARK: - UITableViewDelegate

extension NoKeepEventViewController: UITableViewDelegate, SFSafariViewControllerDelegate {
    
    func tableView(tableView: UITableView, didSelectRowAtIndexPath indexPath:NSIndexPath) {
        
        guard let eventSummaries = eventSummaries else {
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
}
