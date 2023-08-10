import Link from 'next/link';

function Nav() {
    return (
        <div className="bg-gray-800 text-white py-4 px-6">
            <div className="flex justify-between items-center">
                <div>
                    <Link href="/">
                        <p className="mr-6 hover:text-gray-400">Home</p>
                    </Link>
                    <Link href="/about">
                        <p className="hover:text-gray-400">About</p>
                    </Link>
                </div>
                <button className="bg-blue-500 hover:bg-blue-600 py-2 px-4 rounded">
                    Sign Up / Sign In
                </button>
            </div>
        </div>
    );
}

export default Nav;


