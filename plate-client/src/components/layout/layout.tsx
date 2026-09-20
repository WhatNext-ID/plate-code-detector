import { Outlet, useMatches } from 'react-router';
import { AppSidebar } from '@/components/app-sidebar';
import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator,
} from '@/components/ui/breadcrumb';
import { Separator } from '@/components/ui/separator';
import {
  SidebarInset,
  SidebarProvider,
  SidebarTrigger,
} from '@/components/ui/sidebar';
import React from 'react';
import { RouteHandle } from '../../routes/routes';

type RouteMatch = ReturnType<typeof useMatches>[number] & {
  handle?: RouteHandle;
};

export default function Page() {
  const matches = useMatches() as RouteMatch[];

  const breadcrumbs = matches
    .filter((match) => match.handle?.breadcrumb)
    .map((match) => ({
      label: match.handle!.breadcrumb!,
      pathname: match.pathname,
    }));

  return (
    <SidebarProvider>
      <AppSidebar />

      <SidebarInset>
        <header className="flex h-16 shrink-0 items-center gap-2 transition-[width,height] ease-linear group-has-data-[collapsible=icon]/sidebar-wrapper:h-12">
          <div className="flex items-center gap-2 px-4">
            <SidebarTrigger className="-ml-1" />

            <Separator
              orientation="vertical"
              className="mr-2 data-[orientation=vertical]:h-4"
            />

            <Breadcrumb>
              <BreadcrumbList>
                {breadcrumbs.map((breadcrumb, index) => {
                  const isLast = index === breadcrumbs.length - 1;

                  return (
                    <React.Fragment key={breadcrumb.pathname}>
                      {index > 0 && (
                        <BreadcrumbSeparator className="hidden md:block" />
                      )}

                      <BreadcrumbItem
                        className={isLast ? '' : 'hidden md:block'}
                      >
                        {isLast ? (
                          <BreadcrumbPage>{breadcrumb.label}</BreadcrumbPage>
                        ) : (
                          <BreadcrumbLink href={breadcrumb.pathname}>
                            {breadcrumb.label}
                          </BreadcrumbLink>
                        )}
                      </BreadcrumbItem>
                    </React.Fragment>
                  );
                })}
              </BreadcrumbList>
            </Breadcrumb>
          </div>
        </header>

        <div className="ps-2 pe-3 py-3 w-full h-dvh">
          <Outlet />
        </div>
      </SidebarInset>
    </SidebarProvider>
  );
}
