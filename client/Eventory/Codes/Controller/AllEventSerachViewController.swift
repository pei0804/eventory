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

class AllEventSerachViewController: BaseViewController {
    
    var eventSummaries: [EventSummary]? {
        didSet {
            if let eventSummaries = self.eventSummaries where eventSummaries.count == 0 {
                self.tableView.setContentOffset(CGPointZero, animated: false)
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
        
        self.tableView.emptyDataSetSource = self;
        self.tableView.emptyDataSetDelegate = self;
        
        self.tableView.delegate = self
        self.tableView.dataSource = self
        self.freeWordSearchBar.delegate = self
        
        self.tableView.registerNib(UINib(nibName: EventInfoTableViewCellIdentifier, bundle: nil), forCellReuseIdentifier: EventInfoTableViewCellIdentifier)
    }
    
    override func viewWillAppear(animated: Bool) {
        
        super.viewWillAppear(animated)
        eventSummaries = EventManager.sharedInstance.getNewEventAll("")
        
    }
    
    override func didReceiveMemoryWarning() {
        
        super.didReceiveMemoryWarning()
    }

    override func becomeActive(notification: NSNotification) {

        self.refresh()
    }
    
    override func refresh(completed: (() -> Void)? = nil) {
        
        dispatch_async(dispatch_get_main_queue()) {
            let task = [EventManager.sharedInstance.fetchNewEvent()]
            SVProgressHUD.showWithStatus(ServerConnectionMessage)
            Task.all(task).success { _ in
                self.eventSummaries = EventManager.sharedInstance.getNewEventAll("")
                SVProgressHUD.dismiss()
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

// MARK: - UISearchBarDelegate 

extension AllEventSerachViewController: UISearchBarDelegate {
    
    func searchBarCancelButtonClicked(searchBar: UISearchBar) {
        
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

// MARK: - UITableViewDataSource

extension AllEventSerachViewController: UITableViewDataSource {
    
    func numberOfSectionsInTableView(tableView: UITableView) -> Int {
        
        return 1
    }
    
    func tableView(tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        
        if let eventSummaries = self.eventSummaries {
            return eventSummaries.count
        }
        return 0
    }
    
    func tableView(tableView: UITableView, cellForRowAtIndexPath indexPath: NSIndexPath) -> UITableViewCell {
        
        if let cell = self.tableView.dequeueReusableCellWithIdentifier(EventInfoTableViewCellIdentifier, forIndexPath: indexPath) as? EventInfoTableViewCell {
            if let eventSummaries = self.eventSummaries {
                cell.bind(eventSummaries[indexPath.row], viewPageClass: CheckStatus.NoCheck, indexPath: indexPath)
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
}
