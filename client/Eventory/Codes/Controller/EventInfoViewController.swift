//
//  EventInfoViewController.swift
//  Eventory
//
//  Created by jumpei on 2016/08/19.
//  Copyright © 2016年 jumpei. All rights reserved.
//

import UIKit
import SafariServices

class EventInfoViewController: BaseViewController, SFSafariViewControllerDelegate {
    
    var eventSummarys: [EventSummary]? {
        willSet {
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
        let _url:NSURL = NSURL(string: eventSummarys![indexPath.row].url)!
        let _brow = SFSafariViewController(URL: _url, entersReaderIfAvailable: false)
        _brow.delegate = self
        presentViewController(_brow, animated: true, completion: nil)
    }
}