//
//  RegisterPlaceViewController.swift
//  Eventory
//
//  Created by jumpei on 2016/09/05.
//  Copyright © 2016年 jumpei. All rights reserved.
//

import UIKit
import SwiftTask
import SVProgressHUD

class RegisterPlaceViewController: UIViewController {
    
    @IBOutlet weak var tableView: UITableView!
    var checkCount: Int = 0
    var places = [Dictionary<String, AnyObject>]?() {
        didSet {
            tableView.reloadData()
        }
    }
    
    // 設定画面からのアクセスの場合trueになる
    var leftBarButton: UIBarButtonItem = UIBarButtonItem()
    var rightBarButton: UIBarButtonItem = UIBarButtonItem()
    var settingStatus: Bool = false
    
    @IBOutlet weak var searchBar: UISearchBar!
    override func viewDidLoad() {
        
        super.viewDidLoad()
        
        self.searchBar.delegate = self
        self.tableView.delegate = self
        self.tableView.dataSource = self
        
        self.tableView.registerNib(UINib(nibName: CheckListTableViewCellIdentifier, bundle: nil), forCellReuseIdentifier: CheckListTableViewCellIdentifier)
    }
    
    override func viewWillAppear(animated:Bool) {
        
        super.viewWillAppear(animated)
        if self.settingStatus {
            self.leftBarButton = UIBarButtonItem(title: "戻る", style: UIBarButtonItemStyle.Plain, target: self, action: #selector(self.goBack(_:)))
            self.rightBarButton = UIBarButtonItem(title: "適用", style: UIBarButtonItemStyle.Plain, target: self, action: #selector(self.pushSubmitBtn(_:)))
            self.places = UserRegister.sharedInstance.getSettingPlaces()
            self.checkCount = UserRegister.sharedInstance.getUserSettingPlaces().count
        } else {
            self.leftBarButton = UIBarButtonItem(title: "戻る", style: UIBarButtonItemStyle.Plain, target: self, action: #selector(self.goBack(_:)))
            self.rightBarButton = UIBarButtonItem(title: "次へ", style: UIBarButtonItemStyle.Plain, target: self, action: #selector(self.pushSubmitBtn(_:)))
            self.places = EventManager.sharedInstance.placesInitializer()
        }
        self.navigationItem.leftBarButtonItem = self.leftBarButton
        self.navigationItem.rightBarButtonItem = self.rightBarButton
    }
    
    override func viewWillDisappear(animated:Bool) {
        
        super.viewWillDisappear(animated)
    }
    
    override func didReceiveMemoryWarning() {
        
        super.didReceiveMemoryWarning()
    }
    
    @IBAction func goBack(sender: AnyObject) {
        
        self.navigationController?.popToRootViewControllerAnimated(true)

    }
    
    @IBAction func pushEditModeBtn(sender: AnyObject) {
        
        if self.tableView.editing == false {
            self.tableView.editing = true
        } else {
            self.tableView.editing = false
        }
    }
    
    @IBAction func pushSubmitBtn(sender: AnyObject) {
        
        UserRegister.sharedInstance.setUserSettingRegister(self.places, settingClass: SettingClass.Place)
        UserRegister.sharedInstance.setDefaultSettingStatus(true)
        if self.settingStatus {
            self.navigationController?.popToRootViewControllerAnimated(true)
        } else {
            dispatch_async(dispatch_get_main_queue()) {
                SVProgressHUD.showWithStatus(ServerConnectionMessage)
                let task = [EventManager.sharedInstance.fetchNewEvent()]
                Task.all(task).success { _ in
                    SVProgressHUD.dismiss()
                    let storyBoard = UIStoryboard(name: "Main", bundle: nil)
                    let vc: UITabBarController = storyBoard.instantiateViewControllerWithIdentifier("MainMenu") as! UITabBarController
                    self.presentViewController(vc, animated: true, completion: nil)
                    }.failure { _ in
                        SVProgressHUD.dismiss()
                        let alert: UIAlertController = UIAlertController(title: NetworkErrorTitle,message: NetworkErrorMessage, preferredStyle: .Alert)
                        let cancelAction: UIAlertAction = UIAlertAction(title: NetworkErrorButton, style: .Cancel, handler: nil)
                        alert.addAction(cancelAction)
                        self.presentViewController(alert, animated: true, completion: nil)
                }
            }
        }
    }
}


// MARK: - UITableViewDataSource

extension RegisterPlaceViewController: UITableViewDataSource {
    
    func numberOfSectionsInTableView(tableView: UITableView) -> Int {
        
        return 1
    }
    
    func tableView(tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        
        if let places = self.places {
            return places.count
        }
        return 0
    }
    
    func tableView(tableView: UITableView, cellForRowAtIndexPath indexPath: NSIndexPath) -> UITableViewCell {
        
        if let cell = self.tableView.dequeueReusableCellWithIdentifier(CheckListTableViewCellIdentifier, forIndexPath: indexPath) as? CheckListTableViewCell {
            if let places = self.places {
                cell.bind(places[indexPath.row])
                return cell
            }
        }
        return UITableViewCell()
    }
    
    func tableView(tableView: UITableView, commitEditingStyle editingStyle: UITableViewCellEditingStyle, forRowAtIndexPath indexPath: NSIndexPath) {
        
        if editingStyle == UITableViewCellEditingStyle.Delete {
            UserRegister.sharedInstance.deleteSetting(&self.places, index: indexPath.row)
        }
    }
}

// MARK: - UITableViewDelegate

extension RegisterPlaceViewController: UITableViewDelegate {
    func tableView(tableView: UITableView, didSelectRowAtIndexPath indexPath: NSIndexPath) {
        if let cell = self.tableView.cellForRowAtIndexPath(indexPath) as? CheckListTableViewCell {
            cell.checkAction(&self.places, indexPath: indexPath, checkCount: &self.checkCount)
        }
        self.tableView.deselectRowAtIndexPath(indexPath, animated: true)
    }


    func tableView(tableView: UITableView,canEditRowAtIndexPath indexPath: NSIndexPath) -> Bool {
        return true
    }

    func tableView(tableView: UITableView, titleForDeleteConfirmationButtonForRowAtIndexPath indexPath: NSIndexPath) -> String? {
        return "削除"
    }
}


// MARK: - UISearchBarDelegate

extension RegisterPlaceViewController: UISearchBarDelegate {
    
    func searchBarSearchButtonClicked(searchBar: UISearchBar) {
        let text = self.searchBar.text ?? ""
        if !text.isEmpty {
            UserRegister.sharedInstance.insertNewSetting(&self.places, newSetting: text)
            self.searchBar.text = ""
            self.searchBar.resignFirstResponder()
        }
    }
    
    func searchBarCancelButtonClicked(searchBar: UISearchBar) {
        
        self.searchBar.text = ""
        self.searchBar.resignFirstResponder()
    }
}
