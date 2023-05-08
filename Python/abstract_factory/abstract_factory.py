import platform
from abc import ABC, abstractmethod


# Here is an example of how to define
# the abstract factory abstract class that works as interface:
class WidgetFactory(ABC):
    @abstractmethod
    def create_window(self) -> Window:
        pass

    @abstractmethod
    def create_button(self) -> Button:
        pass

    @abstractmethod
    def create_menu(self) -> Menu:
        pass


# We can then define concrete implementations of this interface, such as
class WindowsFactory(WidgetFactory):
    def create_window(self):
        return Window

    def create_button(self):
        return Button

    def create_menu(self):
        return Menu


class MacFactory(WidgetFactory):
    def create_window(self):
        return Window

    def create_button(self):
        return Button

    def create_menu(self):
        return Menu

# In this example, we define two concrete factories,
# WindowsFactory and MacFactory, which implement the WidgetFactory interface.
# Each concrete factory is responsible for creating objects of a specific family.
# For example, the WindowsFactory creates objects with a Windows look and feel,
# while the MacFactory creates objects with a Mac look and feel.

# Then we initialise our factory during runtime.


def main():
    if platform.system() == 'Windows':
        factory = WindowsFactory()
    else:
        factory = MacFactory()

    window = factory.create_window()
    button = factory.create_button()
    menu = factory.create_menu()

    # Use the created objects...
