//
//  NoKeepEventViewController.swift
//  Eventory
//
//  Created by jumpei on 2016/08/22.
//  Copyright © 2016年 jumpei. All rights reserved.
//

import UIKit
import SVProgressHUD
import SwiftTask

class NoKeepEventViewController: BaseTableViewController {

    override func viewDidLoad() {
        super.viewDidLoad()
    }

    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
    }

    override func refresh(completed: (() -> Void)? = nil) {
        dispatch_async(dispatch_get_main_queue()) {
            let task = [EventManager.sharedInstance.fetchNewEvent()]
            Task.all(task).success { _ in
                self.eventSummaries = EventManager.sharedInstance.getNoKeepEventAll()
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
                cell.bind(eventSummaries[indexPath.row], viewPageClass: CheckStatus.NoKeep,indexPath: indexPath)
                return cell
            }
        }
        return UITableViewCell()
    }
}
