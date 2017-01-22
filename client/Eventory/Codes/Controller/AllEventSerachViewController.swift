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
    
    @IBOutlet weak var freeWordSearchBar: UISearchBar!
    
    override func viewDidLoad() {
        
        super.viewDidLoad()
        
//        self.freeWordSearchBar.delegate = self
        self.tableView.registerNib(UINib(nibName: EventInfoTableViewCellIdentifier, bundle: nil), forCellReuseIdentifier: EventInfoTableViewCellIdentifier)
    }
    
    override func refresh(completed: (() -> Void)? = nil) {
        dispatch_async(dispatch_get_main_queue()) {
            let task = [EventManager.sharedInstance.fetchNewEvent()]
            Task.all(task).success { _ in
                self.eventSummaries = EventManager.sharedInstance.getNewEventAll("")
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

    override func tableView(tableView: UITableView, cellForRowAtIndexPath indexPath: NSIndexPath) -> UITableViewCell {

        if let cell = self.tableView.dequeueReusableCellWithIdentifier(EventInfoTableViewCellIdentifier, forIndexPath: indexPath) as? EventInfoTableViewCell {
            if let eventSummaries = self.eventSummaries {
                cell.bind(eventSummaries[indexPath.row], viewPageClass: CheckStatus.NoCheck,indexPath: indexPath)
                return cell
            }
        }
        return UITableViewCell()
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
