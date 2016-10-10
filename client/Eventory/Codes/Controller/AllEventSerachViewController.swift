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

class AllEventSerachViewController: BaseViewController {
    
    var eventSummarys: [EventSummary]? {
        didSet {
            if let eventSummarys = eventSummarys where eventSummarys.count == 0 {
                tableView.setContentOffset(CGPointZero, animated: false)
            }
            self.tableView.reloadData()
        }
    }
    
    @IBOutlet weak var tableView: UITableView!
    @IBOutlet weak var freeWordSearchBar: UISearchBar!
    
    override func viewDidLoad() {
        super.viewDidLoad()
        
        self.scrollView = tableView
        self.addRefreshControl()
        
        self.tableView.emptyDataSetSource = self
        
        tableView.delegate = self
        tableView.dataSource = self
        freeWordSearchBar.delegate = self
        
        self.tableView.registerNib(UINib(nibName: EventInfoTableViewCellIdentifier, bundle: nil), forCellReuseIdentifier: EventInfoTableViewCellIdentifier)
    }
    
    override func viewWillAppear(animated: Bool) {
        
        super.viewWillAppear(animated)
        eventSummarys = EventManager.sharedInstance.getNewEventAll("")
        
    }
    
    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
    }
    
    override func refresh(completed: (() -> Void)? = nil) {
        dispatch_async(dispatch_get_main_queue()) {
            
            let task = [EventManager.sharedInstance.fetchNewEvent()]
            
            Task.all(task).success { _ in
                self.eventSummarys = EventManager.sharedInstance.getNewEventAll("")
                completed?()
                }.failure { _ in
                    let alert: UIAlertController = UIAlertController(title: NetworkErrorTitle,message: NetworkErrorMessage, preferredStyle: .Alert)
                    let cancelAction: UIAlertAction = UIAlertAction(title: NetworkErrorButton, style: .Cancel, handler: nil)
                    alert.addAction(cancelAction)
                    self.presentViewController(alert, animated: true, completion: nil)
                    completed?()
            }
        }
    }
}

extension AllEventSerachViewController: UISearchBarDelegate {
    
    func searchBarCancelButtonClicked(searchBar: UISearchBar) {
        
        freeWordSearchBar.resignFirstResponder()
    }
    
    func searchBarSearchButtonClicked(searchBar: UISearchBar) {
        
        guard let term = freeWordSearchBar.text else {
            freeWordSearchBar.resignFirstResponder()
            return
        }
        eventSummarys = EventManager.sharedInstance.getNewEventAll(term)
        freeWordSearchBar.resignFirstResponder()
    }
    
    override func touchesBegan(touches: Set<UITouch>, withEvent event: UIEvent?) {
        
        if (self.freeWordSearchBar.isFirstResponder()) {
            self.freeWordSearchBar.resignFirstResponder()
        }
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
        return EventInfoCellHeight
    }
}

// MARK: - UITableViewDelegate

extension AllEventSerachViewController: UITableViewDelegate, SFSafariViewControllerDelegate {
    
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
