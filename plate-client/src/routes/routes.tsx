import { createBrowserRouter, Navigate } from 'react-router';

import App from '@/App';
import Landing from '@/pages/LandingPages';
import Ikhtisar from '@/pages/IkhtisarPages';
import Region from '@/pages/region/RegionPages';

export type RouteHandle = {
  breadcrumb?: string;
};

export const router = createBrowserRouter([
  {
    path: '/PlateFrom',
    Component: App,
    handle: {
      breadcrumb: 'Deteksi Plat Nomor',
    },
    children: [
      {
        index: true,
        Component: () => <Navigate to="landing" replace />,
      },
      {
        path: 'landing',
        Component: Landing,
      },
      {
        path: 'ikhtisar',
        Component: Ikhtisar,
        handle: {
          breadcrumb: 'Ikhtisar',
        },
      },
      {
        path: 'region',
        Component: Region,
        handle: {
          breadcrumb: 'Region',
        },
      },
    ],
  },
]);
