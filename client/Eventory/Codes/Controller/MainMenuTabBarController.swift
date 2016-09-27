//
//  MainMenuTabBarController.swift
//  Eventory
//
//  Created by jumpei on 2016/09/11.
//  Copyright © 2016年 jumpei. All rights reserved.
//

import UIKit

class MainMenuTabBarController: UITabBarController {
    
    var newEvent: Int = 0 {
        didSet {
            print("きてます")
            // TODO 最初の時点ではどこのことを指しているかがわからない。
            if newEvent > 0 {
                self.tabBar.items![2].badgeValue = "New"
            } else {
                self.tabBar.items![2].badgeValue = ""
            }
        }
        
    }
    
    let tabBarImages: [String] = ["search", "noKeep", "new", "keep", "setting"]
    
    override func viewDidLoad() {
        super.viewDidLoad()
        self.selectedIndex = 2
        self.tabBar.tintColor = Colors.main
        self.updateBadge(EventManager.sharedInstance.getSelectNewEventAll().count)
        
        guard let items = self.tabBar.items else {
            return
        }
        
        if items.count != tabBarImages.count {
            fatalError("assets tabbarmenu not match count")
        }
        
        for (i,item) in items.enumerate() {
            item.image = UIImage(named: "\(tabBarImages[i])_off.png")!.imageWithRenderingMode(UIImageRenderingMode.AlwaysOriginal)
            item.selectedImage = UIImage(named: "\(tabBarImages[i])_on.png")!.imageWithRenderingMode(UIImageRenderingMode.AlwaysOriginal)
        }
    }
    
    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
        // Dispose of any resources that can be recreated.
    }
    
    override func tabBar(tabBar: UITabBar, didSelectItem item: UITabBarItem) {
        self.updateBadge(EventManager.sharedInstance.getSelectNewEventAll().count)
    }
    
    func updateBadge(newEvent: Int) {
        self.newEvent = newEvent
    }
    
    
    /*
     // MARK: - Navigation
     
     // In a storyboard-based application, you will often want to do a little preparation before navigation
     override func prepareForSegue(segue: UIStoryboardSegue, sender: AnyObject?) {
     // Get the new view controller using segue.destinationViewController.
     // Pass the selected object to the new view controller.
     }
     */
    
}
