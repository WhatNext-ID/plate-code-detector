import { createBrowserRouter, Navigate } from 'react-router';

import App from '@/App';
import Region from '@/pages/region/RegionPages';
import Ikhtisar from '@/pages/IkhtisarPages';
import Landing from '@/pages/LandingPages';

export type RouteHandle = {
  breadcrumb?: string;
  breadcrumbTo?: string;
};

export const router = createBrowserRouter([
  {
    path: '/PlateFrom',

    children: [
      // Landing - no sidebar
      {
        index: true,
        Component: () => <Navigate to="landing" replace />,
      },

      {
        path: 'landing',
        Component: Landing,
      },

      // Application layout
      {
        Component: App,

        handle: {
          breadcrumb: 'Deteksi Plat Nomor',
          breadcrumbTo: '/PlateFrom/landing',
        },

        children: [
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
    ],
  },
]);
