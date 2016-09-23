//
//  RegisterGenreViewController.swift
//  Eventory
//
//  Created by jumpei on 2016/09/05.
//  Copyright © 2016年 jumpei. All rights reserved.
//

import UIKit

class RegisterGenreViewController: UIViewController {
    
    @IBOutlet weak var tableView: UITableView!
    var checkCount: Int = 0
    var genres = [Dictionary<String, AnyObject>]?()
    // 設定画面からのアクセスの場合trueになる
    var settingStatus = false
    
    @IBOutlet weak var searchBar: UISearchBar!
    @IBOutlet weak var submitBtn: UIBarButtonItem!
    
    override func viewDidLoad() {
        
        super.viewDidLoad()
        EventManager.sharedInstance.genreInitializer()
        
        searchBar.delegate = self
        tableView.delegate = self
        tableView.dataSource = self
        
        self.tableView.registerNib(UINib(nibName: CheckListTableViewCellIdentifier, bundle: nil), forCellReuseIdentifier: CheckListTableViewCellIdentifier)
    }
    
    override func viewWillAppear(animated:Bool) {
        
        super.viewWillAppear(animated)
        if settingStatus {
            genres = UserRegister.sharedInstance.getSettingGenres()
            checkCount = UserRegister.sharedInstance.getUserSettingGenres().count
        } else {
            genres = EventManager.sharedInstance.genreInitializer()
        }
        self.tableView.reloadData()
    }
    
    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
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
//            let alert: UIAlertController = UIAlertController(title: "最低１つ選んでください。", message: "１つも選択されていないと検索できません。", preferredStyle: .Alert)
//            let cancelAction: UIAlertAction = UIAlertAction(title: "OK", style: .Cancel, handler: nil)
//            alert.addAction(cancelAction)
//            self.presentViewController(alert, animated: true, completion: nil)
//            return
//        }
        
        UserRegister.sharedInstance.setUserSettingRegister(genres, settingClass: SettingClass.Genre)
        
        if settingStatus {
            navigationController?.popToRootViewControllerAnimated(true)
        } else {
            let vc = UIStoryboard(name:"Register", bundle: nil).instantiateViewControllerWithIdentifier(RegisterPlaceViewControllerIdentifier)
            self.navigationController?.pushViewController(vc, animated: true)
        }
    }
}

// MARK: - UITableViewDataSource

extension RegisterGenreViewController: UITableViewDataSource {
    
    func numberOfSectionsInTableView(tableView: UITableView) -> Int {
        
        return 1
    }
    
    func tableView(tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        
        if let genres = genres {
            return genres.count
        }
        return 0
    }
    
    func tableView(tableView: UITableView, cellForRowAtIndexPath indexPath: NSIndexPath) -> UITableViewCell {
        
        if let cell = tableView.dequeueReusableCellWithIdentifier(CheckListTableViewCellIdentifier, forIndexPath: indexPath) as? CheckListTableViewCell {
            if let genres = genres {
                cell.bind(genres[indexPath.row])
                return cell
            }
        }
        return UITableViewCell()
    }
    
    func tableView(tableView: UITableView, commitEditingStyle editingStyle: UITableViewCellEditingStyle, forRowAtIndexPath indexPath: NSIndexPath) {
        if editingStyle == UITableViewCellEditingStyle.Delete {
            UserRegister.sharedInstance.deleteSetting(&genres, index: indexPath.row)
            tableView.deleteRowsAtIndexPaths([indexPath], withRowAnimation: UITableViewRowAnimation.Automatic)
        }
    }
}

// MARK: - UITableViewDelegate

extension RegisterGenreViewController: UITableViewDelegate {
    
    func tableView(tableView: UITableView, didSelectRowAtIndexPath indexPath: NSIndexPath) {
        if let cell = tableView.cellForRowAtIndexPath(indexPath) as? CheckListTableViewCell {
            cell.checkAction(&genres, indexPath: indexPath, checkCount: &checkCount)
        }
        tableView.deselectRowAtIndexPath(indexPath, animated: true)
    }
}

// MARK: - UISearchBarDelegate

extension RegisterGenreViewController: UISearchBarDelegate {
    
    func searchBarSearchButtonClicked(searchBar: UISearchBar) {
        let text = searchBar.text ?? ""
        if !text.isEmpty {
            UserRegister.sharedInstance.insertNewSetting(&genres, newSetting: String(searchBar.text!))
            searchBar.text = ""
            searchBar.resignFirstResponder()
        }
    }
    
    func searchBarCancelButtonClicked(searchBar: UISearchBar) {
        
        searchBar.text = ""
        searchBar.resignFirstResponder()
    }
}