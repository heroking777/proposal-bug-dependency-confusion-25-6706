import pkg_resources

def check_dependency_confusion(package_name):
    try:
        # Attempt to get the distribution of the package
        dist = pkg_resources.get_distribution(package_name)
        print(f"Package {package_name} is installed with version {dist.version}")
    except pkg_resources.DistributionNotFound:
        print(f"Package {package_name} is not installed")

# Example usage
check_dependency_confusion("some-package")