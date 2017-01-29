//
//  define.swift
//  Eventory
//
//  Created by jumpei on 2016/09/06.
//  Copyright © 2016年 jumpei. All rights reserved.
//

import Foundation
import UIKit

// Cell
let EventInfoTableViewCellIdentifier: String = "EventInfoTableViewCell"
let CheckListTableViewCellIdentifier: String = "CheckListTableViewCell"
let SettingTableViewCellIdentifier: String   = "SettingTableViewCell"

// Controller
let RegisterPlaceViewControllerIdentifier: String   = "RegisterPlaceViewController"
let RegisterGenreViewControllerIdentifier: String   = "RegisterGenreViewController"
let SettingViewControllerIdentifier: String         = "SettingViewController"

// Alert
let NetworkErrorTitle: String   = "サーバー接続に失敗しました"
let NetworkErrorMessage: String = "端末がインターネットが使えない。またはサーバーに問題がありました。"
let NetworkErrorButton: String  = "確認"
let ServerConnectionMessage: String = "サーバーと通信中"

// Cell
let EventInfoCellHeight = CGFloat(250)

enum ApiId: Int {
    
    case Atdn       = 0
    case Connpass   = 1
    case Doorkeeper = 2
    case None       = 5
    
    func getName() -> String {
        
        switch self {
        case .Atdn:       return "ATDN"
        case .Connpass:   return "Connpass"
        case .Doorkeeper: return "Doorkeeper"
        case .None:       return ""
        }
    }
}

// UserDefault

enum SettingClass: Int {
    
    case Genre      = 0
    case Place      = 1
    case Status     = 2
    case None       = 5
    
    func getUserSettingKey() -> String {
        
        switch self {
        case .Genre:    return "RegisterGenres"
        case .Place:    return "RegisterPlaces"
        case .Status:   return ""
        case .None:     return ""
        }
    }
    
    func getSettingKey() -> String {
        
        switch self {
        case .Genre:    return "SettingGenres"
        case .Place:    return "SettingPlaces"
        case .Status:   return "SettingStatus"
        case .None:     return ""
        }
    }
}

let userEventInfoUpdatedAt = "UserEventInfoUpdatedAt"
