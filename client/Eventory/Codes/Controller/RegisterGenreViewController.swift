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
    var genres = [Dictionary<String, AnyObject>]?() {
        didSet {
            self.tableView.reloadData()
        }
    }
    // 設定画面からのアクセスの場合trueになる
    var leftBarButton: UIBarButtonItem = UIBarButtonItem()
    var rightBarButton: UIBarButtonItem = UIBarButtonItem()
    var settingStatus: Bool = false
    
    @IBOutlet weak var searchBar: UISearchBar!
    @IBOutlet weak var submitBtn: UIBarButtonItem!
    
    override func viewDidLoad() {
        
        super.viewDidLoad()
        EventManager.sharedInstance.genreInitializer()
        
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
            self.genres = UserRegister.sharedInstance.getSettingGenres()
            self.checkCount = UserRegister.sharedInstance.getUserSettingGenres().count
        } else {
            self.rightBarButton = UIBarButtonItem(title: "次へ", style: UIBarButtonItemStyle.Plain, target: self, action: #selector(self.pushSubmitBtn(_:)))
            self.genres = EventManager.sharedInstance.genreInitializer()
        }
        self.navigationItem.leftBarButtonItem = self.leftBarButton
        self.navigationItem.rightBarButtonItem = self.rightBarButton
    }
    
    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
    }
    
    @IBAction func pushEditModeBtn(sender: AnyObject) {
        if self.tableView.editing == false {
            self.tableView.editing = true
        } else {
            self.tableView.editing = false
        }
    }
    
    @IBAction func goBack(sender: AnyObject) {
        
        self.navigationController?.popToRootViewControllerAnimated(true)
        
    }
    
    @IBAction func pushSubmitBtn(sender: AnyObject) {
        UserRegister.sharedInstance.setUserSettingRegister(self.genres, settingClass: SettingClass.Genre)
        if self.settingStatus {
            self.navigationController?.popToRootViewControllerAnimated(true)
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
        if let genres = self.genres {
            return genres.count
        }
        return 0
    }
    
    func tableView(tableView: UITableView, cellForRowAtIndexPath indexPath: NSIndexPath) -> UITableViewCell {
        if let cell = self.tableView.dequeueReusableCellWithIdentifier(CheckListTableViewCellIdentifier, forIndexPath: indexPath) as? CheckListTableViewCell {
            if let genres = self.genres {
                cell.bind(genres[indexPath.row])
                return cell
            }
        }
        return UITableViewCell()
    }
    
    func tableView(tableView: UITableView, commitEditingStyle editingStyle: UITableViewCellEditingStyle, forRowAtIndexPath indexPath: NSIndexPath) {
        if editingStyle == UITableViewCellEditingStyle.Delete {
            UserRegister.sharedInstance.deleteSetting(&self.genres, index: indexPath.row)
        }
    }
}

// MARK: - UITableViewDelegate

extension RegisterGenreViewController: UITableViewDelegate {
    
    func tableView(tableView: UITableView, didSelectRowAtIndexPath indexPath: NSIndexPath) {
        if let cell = self.tableView.cellForRowAtIndexPath(indexPath) as? CheckListTableViewCell {
            cell.checkAction(&self.genres, indexPath: indexPath, checkCount: &self.checkCount)
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

extension RegisterGenreViewController: UISearchBarDelegate {
    
    func searchBarSearchButtonClicked(searchBar: UISearchBar) {
        let text = self.searchBar.text ?? ""
        if !text.isEmpty {
            UserRegister.sharedInstance.insertNewSetting(&self.genres, newSetting: text)
            self.searchBar.text = ""
            self.searchBar.resignFirstResponder()
        }
    }
    
    func searchBarCancelButtonClicked(searchBar: UISearchBar) {
        self.searchBar.text = ""
        self.searchBar.resignFirstResponder()
    }
}
