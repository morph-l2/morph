import devnet
from devnet.setup_nodes import setup_devnet_nodes


def main():
    setup_devnet_nodes()
    devnet.main()


if __name__ == '__main__':
    main()