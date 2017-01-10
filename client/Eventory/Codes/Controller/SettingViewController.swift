//
//  SettingViewController.swift
//  Eventory
//
//  Created by jumpei on 2016/09/10.
//  Copyright © 2016年 jumpei. All rights reserved.
//

import UIKit

class SettingViewController: UIViewController {
    
    @IBOutlet weak var tableView: UITableView!
    let settingClasses: [Dictionary<String, String>] = [
        [
            "name": "興味のあるジャンル",
            "controller": RegisterGenreViewControllerIdentifier
        ],
        [
            "name": "開催地",
            "controller": RegisterPlaceViewControllerIdentifier
        ]
    ]
    override func viewDidLoad() {
        super.viewDidLoad()
        self.tableView.dataSource = self
        self.tableView.delegate = self
        self.tableView.separatorInset = UIEdgeInsetsZero
        self.tableView.layoutMargins = UIEdgeInsetsZero
    }
    
    override func viewWillAppear(animated:Bool) {
        
        super.viewWillAppear(animated)
    }
    
    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
    }
}

// MARK: - UITableViewDataSource

extension SettingViewController: UITableViewDataSource {
    
    func numberOfSectionsInTableView(tableView: UITableView) -> Int {
        
        return 1
    }
    
    func tableView(tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        
        return settingClasses.count
    }
    
    func tableView(tableView: UITableView, cellForRowAtIndexPath indexPath: NSIndexPath) -> UITableViewCell {
        let cell = tableView.dequeueReusableCellWithIdentifier(SettingTableViewCellIdentifier, forIndexPath: indexPath)
        cell.textLabel?.text = settingClasses[indexPath.row]["name"]
        cell.accessoryType = .DisclosureIndicator
        return cell
    }
}


// MARK: - UITableViewDelegate

extension SettingViewController: UITableViewDelegate {
    
    func tableView(tableView: UITableView, didSelectRowAtIndexPath indexPath: NSIndexPath) {
        
        if let nextVc = settingClasses[indexPath.row]["controller"] {
           
            // どのコントローラに遷移するか判定している
            // 既存にあるRegister系は少し煩雑になっているが、初期登録アクセスかを判定しているだけ
            if nextVc == RegisterPlaceViewControllerIdentifier {
                let vc = UIStoryboard(name:"Register", bundle: nil).instantiateViewControllerWithIdentifier(nextVc) as! RegisterPlaceViewController
                vc.settingStatus = true
                self.navigationController?.pushViewController(vc, animated: true)
            } else if nextVc == RegisterGenreViewControllerIdentifier {
                let vc = UIStoryboard(name:"Register", bundle: nil).instantiateViewControllerWithIdentifier(nextVc) as! RegisterGenreViewController
                vc.settingStatus = true
                self.navigationController?.pushViewController(vc, animated: true)
            }
        }
    }
}

