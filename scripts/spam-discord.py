import pyautogui as pag
import keyboard
import time

def main():
  commandToBeTested = '/pickgame'
  count = 0

  print("[script] The script is going to start.\nGo to Lau's Discord server screen. You have 5 seconds.")
  time.sleep(5)

  while keyboard.is_pressed('q') != True:
    count += 1

    p = pag.pixel(700, 810)

    if p[0] == 56 and p[1] == 58 and p[2] == 64:
      pag.click(700, 810)
      pag.write(commandToBeTested, 0.1)
      pag.press('enter')
      pag.press('enter')

    time.sleep(3)

  print(f'[script] Performed {commandToBeTested} {count} times.')

if __name__ == "__main__":
  main()
