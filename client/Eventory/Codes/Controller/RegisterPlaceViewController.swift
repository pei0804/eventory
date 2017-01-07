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
        
        searchBar.delegate = self
        tableView.delegate = self
        tableView.dataSource = self
        
        self.tableView.registerNib(UINib(nibName: CheckListTableViewCellIdentifier, bundle: nil), forCellReuseIdentifier: CheckListTableViewCellIdentifier)
    }
    
    override func viewWillAppear(animated:Bool) {
        
        super.viewWillAppear(animated)
        
        if settingStatus {
            leftBarButton = UIBarButtonItem(title: "設定", style: UIBarButtonItemStyle.Plain, target: self, action: #selector(self.pushSubmitBtn(_:)))
            rightBarButton = UIBarButtonItem(title: "編集", style: UIBarButtonItemStyle.Plain, target: self, action: #selector(self.pushEditModeBtn(_:)))
            places = UserRegister.sharedInstance.getSettingPlaces()
            checkCount = UserRegister.sharedInstance.getUserSettingPlaces().count
        } else {
            leftBarButton = UIBarButtonItem(title: "戻る", style: UIBarButtonItemStyle.Plain, target: self, action: #selector(self.goBack(_:)))
            rightBarButton = UIBarButtonItem(title: "次へ", style: UIBarButtonItemStyle.Plain, target: self, action: #selector(self.pushSubmitBtn(_:)))
            places = EventManager.sharedInstance.placesInitializer()
        }
        self.navigationItem.leftBarButtonItem = leftBarButton
        self.navigationItem.rightBarButtonItem = rightBarButton
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
        if tableView.editing == false {
            tableView.editing = true
        } else {
            tableView.editing = false
        }
    }
    
    @IBAction func pushSubmitBtn(sender: AnyObject) {
        
        // あえて、ひとつも選ばない人もいる可能性があるので許容してみる
        //        if checkCount <= 0 {
        //            let alert: UIAlertController = UIAlertController(title: "最低１つ選んでください。", message: "１つも選択されていないと有効な結果が得られません。", preferredStyle: .Alert)
        //            let cancelAction: UIAlertAction = UIAlertAction(title: "OK", style: .Cancel, handler: nil)
        //            alert.addAction(cancelAction)
        //            self.presentViewController(alert, animated: true, completion: nil)
        //            return
        //        }
        
        UserRegister.sharedInstance.setUserSettingRegister(places, settingClass: SettingClass.Place)
        UserRegister.sharedInstance.setDefaultSettingStatus(true)
        
        if settingStatus {
            navigationController?.popToRootViewControllerAnimated(true)
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
                        // TODOなんかする
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
        
        if let places = places {
            return places.count
        }
        return 0
    }
    
    func tableView(tableView: UITableView, cellForRowAtIndexPath indexPath: NSIndexPath) -> UITableViewCell {
        
        if let cell = tableView.dequeueReusableCellWithIdentifier(CheckListTableViewCellIdentifier, forIndexPath: indexPath) as? CheckListTableViewCell {
            if let places = places {
                cell.bind(places[indexPath.row])
                return cell
            }
        }
        return UITableViewCell()
    }
    
    func tableView(tableView: UITableView, commitEditingStyle editingStyle: UITableViewCellEditingStyle, forRowAtIndexPath indexPath: NSIndexPath) {
        if editingStyle == UITableViewCellEditingStyle.Delete {
            UserRegister.sharedInstance.deleteSetting(&places, index: indexPath.row)
        }
    }
}

// MARK: - UITableViewDelegate

extension RegisterPlaceViewController: UITableViewDelegate {
    func tableView(tableView: UITableView, didSelectRowAtIndexPath indexPath: NSIndexPath) {
        if let cell = tableView.cellForRowAtIndexPath(indexPath) as? CheckListTableViewCell {
            cell.checkAction(&places, indexPath: indexPath, checkCount: &checkCount)
        }
        tableView.deselectRowAtIndexPath(indexPath, animated: true)
    }
}


// MARK: - UISearchBarDelegate

extension RegisterPlaceViewController: UISearchBarDelegate {
    
    func searchBarSearchButtonClicked(searchBar: UISearchBar) {
        let text = searchBar.text ?? ""
        if !text.isEmpty {
            UserRegister.sharedInstance.insertNewSetting(&places, newSetting: text)
            searchBar.text = ""
            searchBar.resignFirstResponder()
        }
    }
    
    func searchBarCancelButtonClicked(searchBar: UISearchBar) {
        
        searchBar.text = ""
        searchBar.resignFirstResponder()
    }
}
