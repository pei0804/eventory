//
//  AllEventSerachViewController.swift
//  Eventory
//
//  Created by jumpei on 2016/09/18.
//  Copyright © 2016年 jumpei. All rights reserved.
//

import UIKit
import SafariServices

class AllEventSerachViewController: BaseViewController, SFSafariViewControllerDelegate {
    
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
    }
    
    override func viewWillAppear(animated: Bool) {
        
        super.viewWillAppear(animated)
        eventSummarys = EventManager.sharedInstance.getNewEventAll()
        
    }
    
    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
    }
    
    override func refresh(completed: (() -> Void)? = nil) {
        eventSummarys = EventManager.sharedInstance.getNewEventAll()
        completed?()
    }
}

extension AllEventSerachViewController: UITableViewDataSource {
    
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

extension AllEventSerachViewController: UITableViewDelegate {
    
    func tableView(tableView: UITableView, didSelectRowAtIndexPath indexPath:NSIndexPath) {
        let _url:NSURL = NSURL(string: eventSummarys![indexPath.row].url)!
        let _brow = SFSafariViewController(URL: _url, entersReaderIfAvailable: false)
        _brow.delegate = self
        presentViewController(_brow, animated: true, completion: nil)
    }
}
