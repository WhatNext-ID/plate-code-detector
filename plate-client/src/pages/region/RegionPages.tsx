import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table';
import { RegionList } from '@/types/TableInterfaces';
import { listRegionPlate } from '@/utils/network';
import { useQuery } from '@tanstack/react-query';
import {
  ColumnDef,
  flexRender,
  useReactTable,
  getCoreRowModel,
} from '@tanstack/react-table';

export default function Region() {
  const { data, isLoading } = useQuery({
    queryKey: ['list-region'],
    queryFn: async () => listRegionPlate(),
  });

  const columns: ColumnDef<RegionList>[] = [
    {
      header: 'No.',
      cell: ({ row }) => row.index + 1,
    },
    {
      accessorKey: 'regionArea',
      header: 'Daerah',
    },
    {
      accessorKey: 'regionCode',
      header: 'Kode Plat',
    },
    {
      accessorKey: 'regionNote',
      header: 'Keterangan',
    },
  ];

  const table = useReactTable({
    data: data ?? [],
    columns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <div>
      <Table>
        <TableCaption>Daftar Kode Plat Kendaraan Bagian Depan</TableCaption>
        <TableHeader>
          {table.getHeaderGroups().map((headerGroup) => (
            <TableRow key={headerGroup.id}>
              {headerGroup.headers.map((header) => (
                <TableHead key={header.id}>
                  {header.isPlaceholder
                    ? null
                    : flexRender(
                        header.column.columnDef.header,
                        header.getContext()
                      )}
                </TableHead>
              ))}
            </TableRow>
          ))}
        </TableHeader>
        <TableBody>
          {table.getRowModel().rows.length ? (
            table.getRowModel().rows.map((row) => (
              <TableRow key={row.id}>
                {row.getVisibleCells().map((cell) => (
                  <TableCell key={cell.id}>
                    {flexRender(cell.column.columnDef.cell, cell.getContext())}
                  </TableCell>
                ))}
              </TableRow>
            ))
          ) : (
            <TableRow>
              <TableCell colSpan={columns.length} className="text-center">
                {isLoading ? 'Loading...' : 'Tidak ada data'}
              </TableCell>
            </TableRow>
          )}
        </TableBody>
      </Table>
    </div>
  );
}
