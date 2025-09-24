import { createBrowserRouter, Navigate } from 'react-router';
import App from '@/App';
import Dashboard from '@/pages/DashboardPages';
import Region from '@/pages/region/RegionPages';

export const router = createBrowserRouter([
  {
    path: '/PlateFrom',
    Component: App,
    children: [
      {
        index: true,
        Component: () => <Navigate to="ikhtisar" replace />,
      },
      {
        path: 'ikhtisar',
        Component: Dashboard,
      },
      {
        path: 'region',
        Component: Region,
      },
    ],
  },
]);
