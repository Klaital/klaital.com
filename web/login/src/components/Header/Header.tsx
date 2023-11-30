import "./Header.css"
export function Header() {
    return <>
        <header className="site-header">
            <div className="wrapper site-header_wrapper">
                <div className="header-logo">
                    <h1><a className="header-tile header-title" tabIndex={0} href="/">Klaital.com Login</a></h1>
                </div>
                <div className="header-menu">
                    <a className="header-tile" href="/">Home</a>
                    <a className="header-tile" href="/preferences">Account</a>
                    <a className="header-tile" href="/login">Login</a>
                    <a className="header-tile" href="/register">Register</a>
                </div>
            </div>
        </header>
    </>
}
