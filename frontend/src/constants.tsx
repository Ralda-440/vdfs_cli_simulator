import { Icon } from '@iconify/react';

import { SideNavItem } from './types';

export const SIDENAV_ITEMS: SideNavItem[] = [
  {
    title: 'Cli',
    path: '/',
    icon: <Icon icon="fa6-solid:terminal" width="24" height="24" />,
  },
  /*
  {
    title: 'Projects',
    path: '/projects',
    icon: <Icon icon="lucide:folder" width="24" height="24" />,
    submenu: true,
    subMenuItems: [
      { title: 'All', path: '/projects' },
      { title: 'Web Design', path: '/projects/web-design' },
      { title: 'Graphic Design', path: '/projects/graphic-design' },
    ],
  },
  */
  {
    title: 'Explorer',
    path: '/explorer',
    icon: <Icon icon="octicon:file-directory-16" width="24" height="24" />,
  },
  /*
  {
    title: 'Settings',
    path: '/settings',
    icon: <Icon icon="lucide:settings" width="24" height="24" />,
    submenu: true,
    subMenuItems: [
      { title: 'Account', path: '/settings/account' },
      { title: 'Privacy', path: '/settings/privacy' },
    ],
  },
  */
  {
    title: 'Reportes',
    path: '/reportes',
    icon: <Icon icon="fluent-mdl2:report-library-mirrored" width="24" height="24" />,
  },
];
