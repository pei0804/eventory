//
//  PushNotificationManager.swift
//  Eventory
//
//  Created by jumpei on 2016/09/11.
//  Copyright © 2016年 jumpei. All rights reserved.
//

import UIKit
import Firebase

class PushNotificationManager {
    
    static let sharedInstance = PushNotificationManager()
    
    var enabled: Bool {
        if let currentSettings = UIApplication.sharedApplication().currentUserNotificationSettings() {
            if currentSettings.types != .None {
                return true
            }
        }
        return false
    }
    
    // TODO:　プッシュ通知の実装
    func registerRemote() {
        let application = UIApplication.sharedApplication()
        let notificationSettings = UIUserNotificationSettings(
            forTypes: [.Badge, .Sound, .Alert], categories: nil)
        application.registerUserNotificationSettings(notificationSettings)
        application.registerForRemoteNotifications()
        FIRApp.configure()
    }
}