//
//  define.swift
//  Eventory
//
//  Created by jumpei on 2016/09/06.
//  Copyright © 2016年 jumpei. All rights reserved.
//

import Foundation


// Cell
let EventInfoTableViewCellIdentifier = "EventInfoTableViewCell"
let CheckListTableViewCellIdentifier = "CheckListTableViewCell"
let SettingTableViewCellIdentifier   = "SettingTableViewCell"

// Controller
let RegisterPlaceViewControllerIdentifier   = "RegisterPlaceViewController"
let RegisterGenreViewControllerIdentifier   = "RegisterGenreViewController"
let SettingViewControllerIdentifier         = "SettingViewController"

enum SettingClass: Int {
    case Genre = 0
    case Place = 1
    case None = 5
    
    func getUserSettingKey() -> String {
        
        switch self {
        case .Genre:    return "RegisterGenres"
        case .Place:    return "RegisterPlaces"
        case .None:     return ""
        }
    }
    
    func getSettingKey() -> String {
        
        switch self {
        case .Genre:    return "SettingGenres"
        case .Place:    return "SettingPlaces"
        case .None:     return ""
        }
    }
}