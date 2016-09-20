//
//  MainMenuTabBarController.swift
//  Eventory
//
//  Created by jumpei on 2016/09/11.
//  Copyright © 2016年 jumpei. All rights reserved.
//

import UIKit

class MainMenuTabBarController: UITabBarController {
    
    let tabBarImages: [String] = ["search", "noKeep", "new", "keep", "setting"]
    
    override func viewDidLoad() {
        super.viewDidLoad()
        self.selectedIndex = 2
        self.tabBar.tintColor = Colors.mainColor
        
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
    
    
    /*
     // MARK: - Navigation
     
     // In a storyboard-based application, you will often want to do a little preparation before navigation
     override func prepareForSegue(segue: UIStoryboardSegue, sender: AnyObject?) {
     // Get the new view controller using segue.destinationViewController.
     // Pass the selected object to the new view controller.
     }
     */
    
}
