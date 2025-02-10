import os
import time
import pyautogui
from playwright.sync_api import sync_playwright
from dotenv import load_dotenv
with sync_playwright() as p:
    browser = p.chromium.launch(headless=False)
    page = browser.new_page()
    page.goto("https://cmsweb.cms.csub.edu/psp/CBAKPRD/?cmd=login")
    load_dotenv()
    email = os.getenv("EMAIL")
    page.fill('input[name="loginfmt"]', email)
    page.wait_for_selector('#idSIButton9')  # Using the ID to find the button
    page.click('#idSIButton9')
    passwd = os.getenv("PASSWD")
    page.fill('#i0118',passwd)
    page.wait_for_selector('#idSIButton9')  # Using the ID to find the button
    page.click('#idSIButton9')
    time.sleep(5)
    pin = os.getenv("PIN")
    pyautogui.write(pin)
    page.wait_for_selector('#trust-browser-button')
    page.click('#trust-browser-button')
    page.wait_for_selector('#idSIButton9')  # Using the ID to find the button
    page.click('#idSIButton9')
    page.wait_for_timeout(1000000)

