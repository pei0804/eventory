//
//  AppDelegate.swift
//  Eventory
//
//  Created by jumpei on 2016/08/19.
//  Copyright © 2016年 jumpei. All rights reserved.
//

import UIKit

@UIApplicationMain
class AppDelegate: UIResponder, UIApplicationDelegate {
    
    var window: UIWindow?
    
    var newEvent: Int = 0 {
        willSet(newEvent) {
            if newEvent <= 0 {
                if let tabBarController = self.window?.rootViewController as? UITabBarController {
                    tabBarController.tabBar.items![2].badgeValue = nil
                }
            } else {
                if let tabBarController = self.window?.rootViewController as? UITabBarController {
                    tabBarController.tabBar.items![2].badgeValue = "New"
                }
            }
        }
        
    }
    
    func application(application: UIApplication, didFinishLaunchingWithOptions launchOptions: [NSObject: AnyObject]?) -> Bool {
        
        EventManager.sharedInstance.eventInitializer()
        self.newEvent = EventManager.sharedInstance.getSelectNewEventAll().count

        
        if UserRegister.sharedInstance.getUserSettingGenres().isEmpty || UserRegister.sharedInstance.getUserSettingPlaces().isEmpty {
            let storyboard = UIStoryboard(name: "Register", bundle: nil)
            let initialViewController = storyboard.instantiateViewControllerWithIdentifier("signUp")
            self.window?.rootViewController = initialViewController
            self.window?.makeKeyAndVisible()
        } else {
            let storyboard = UIStoryboard(name: "Main", bundle: nil)
            let initialViewController = storyboard.instantiateViewControllerWithIdentifier("MainMenu")
            self.window?.rootViewController = initialViewController
            self.window?.makeKeyAndVisible()
        }
        return true
    }
    
    func applicationWillResignActive(application: UIApplication) {
        // Sent when the application is about to move from active to inactive state. This can occur for certain types of temporary interruptions (such as an incoming phone call or SMS message) or when the user quits the application and it begins the transition to the background state.
        // Use this method to pause ongoing tasks, disable timers, and throttle down OpenGL ES frame rates. Games should use this method to pause the game.
    }
    
    func applicationDidEnterBackground(application: UIApplication) {
        // Use this method to release shared resources, save user data, invalidate timers, and store enough application state information to restore your application to its current state in case it is terminated later.
        // If your application supports background execution, this method is called instead of applicationWillTerminate: when the user quits.
    }
    
    func applicationWillEnterForeground(application: UIApplication) {
        // Called as part of the transition from the background to the inactive state; here you can undo many of the changes made on entering the background.
    }
    
    func applicationDidBecomeActive(application: UIApplication) {
        // Restart any tasks that were paused (or not yet started) while the application was inactive. If the application was previously in the background, optionally refresh the user interface.
    }
    
    func applicationWillTerminate(application: UIApplication) {
        // Called when the application is about to terminate. Save data if appropriate. See also applicationDidEnterBackground:.
    }
    
    func application(application: UIApplication, performFetchWithCompletionHandler completionHandler: (UIBackgroundFetchResult) -> Void) {
    }
    
    
}

