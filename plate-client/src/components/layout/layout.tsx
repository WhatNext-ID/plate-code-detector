import { Outlet } from 'react-router';
import { SidebarProvider, SidebarTrigger } from '../ui/sidebar';
import { AppSidebar } from './app-sidebar';

export default function Layout() {
  return (
    <SidebarProvider>
      <AppSidebar />
      <main>
        <SidebarTrigger />
        <div className="ps-2 pe-3 py-3 w-full h-dvh">
          <Outlet />
        </div>
      </main>
    </SidebarProvider>
  );
}
