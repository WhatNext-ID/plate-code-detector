import { LayoutDashboardIcon, MapPinned } from 'lucide-react';

import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from '@/components/ui/sidebar';
import { NavLink } from 'react-router';

// Menu items.
const items = [
  {
    title: 'Home',
    url: 'ikhtisar',
    icon: LayoutDashboardIcon,
  },
  {
    title: 'Region',
    url: 'region',
    icon: MapPinned,
  },
];

export function AppSidebar() {
  return (
    <Sidebar collapsible="icon" variant="floating">
      <SidebarContent>
        <SidebarGroup>
          <SidebarGroupLabel className="text-md font-black">
            PlateFrom
          </SidebarGroupLabel>
          <SidebarGroupContent>
            <SidebarMenu>
              {items.map((item) => (
                <SidebarMenuItem key={item.title}>
                  <NavLink to={item.url}>
                    {({ isActive }) => (
                      <SidebarMenuButton asChild isActive={isActive}>
                        <span>
                          <item.icon
                            className={isActive ? 'text-red-600' : ''}
                          />
                          <span
                            className={isActive ? 'font-bold text-red-600' : ''}
                          >
                            {item.title}
                          </span>
                        </span>
                      </SidebarMenuButton>
                    )}
                  </NavLink>
                </SidebarMenuItem>
              ))}
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>
    </Sidebar>
  );
}
