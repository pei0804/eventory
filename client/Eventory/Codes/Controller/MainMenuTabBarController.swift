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
            if newEvent > 0 {
                self.tabBar.items![2].badgeValue = "New"
            } else {
                self.tabBar.items![2].badgeValue = nil
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
}
