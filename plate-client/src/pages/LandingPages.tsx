import { useNavigate } from 'react-router';

export default function Landing() {
  const navigate = useNavigate();

  return (
    <main className="min-h-screen w-full p-6">
      <p>
        Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
        tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim
        veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea
        commodo consequat. Duis aute irure dolor in reprehenderit in voluptate
        velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint
        occaecat cupidatat non proident, sunt in culpa qui officia deserunt
        mollit anim id est laborum.
      </p>

      <button
        type="button"
        onClick={() => navigate('/PlateFrom/ikhtisar')}
        className="
          cursor-pointer
          rounded-md
          border
          border-primary
          bg-primary
          px-4
          py-2
          text-primary-foreground
          transition-colors
          hover:border-border
          hover:bg-transparent
          hover:text-foreground
        "
      >
        Go to Ikhtisar
      </button>
    </main>
  );
}
