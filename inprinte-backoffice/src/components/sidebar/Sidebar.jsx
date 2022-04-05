import './sidebar.css';
import HomeIcon from '@material-ui/icons/Home';
import React from 'react'
import AssessmentIcon from '@material-ui/icons/Assessment';

export default function Sidebar() {
  return (
    <div className="sidebar">
        <div className="siderbarWrapper">
            <div className="sidebarMenu">
                <div className="sidebarTitle">
                    <h3>Dashboard 🧙‍♂️</h3>
                    <ul className="sidebarList">
                        <li className="sidebarListItem">
                        <a href="/">
                            <HomeIcon className="sidebarIcon" />
                            Home
                        </a>
                        </li>
                        <li className="sidebarListItem">
                        <a href="/analytics">
                            <AssessmentIcon className="sidebarIcon" />
                            Analytics
                        </a>
                        </li>
                    </ul>
                </div>
            </div>
        </div>

    </div>
  )
}
