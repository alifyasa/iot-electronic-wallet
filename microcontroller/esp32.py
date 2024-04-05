import machine
import time
import network
import urequests

#WiFi connection information
ssid = 'Capillary2944'
password = 'kelompok'

# Connect to WiFi
def connect_to_wifi():
    sta_if = network.WLAN(network.STA_IF)
    if not sta_if.isconnected():
        print('Connecting to network...')
        sta_if.active(True)
        sta_if.connect(ssid, password)
        while not sta_if.isconnected():
            pass
        print('Network connected!')

# Base frequency
def button_press(pin):
    print('Paying...')
    post_request()
    
def led_fail():
    for i in range(10):
        led.value(1)  # Turn LED on
        time.sleep(0.25)
        led.value(0)  # Turn LED off
        time.sleep(0.25)

def led_success():
    led.value(1)
    time.sleep(5)
    led.value(0)
    time.sleep(1)

def post_request():
    response = urequests.post("http://192.168.130.90:4000/pay")
    new_balance = response.text
    status_code = response.status_code
    
    if status_code == 200:
        led_success()
    else:
        led_fail()
    
    wallet_balance = new_balance
    print(new_balance)
    print(status_code)
    response.close()
    
boot_button = machine.Pin(0, machine.Pin.IN, machine.Pin.PULL_UP)
boot_button.irq(trigger=machine.Pin.IRQ_FALLING, handler=button_press)
led = machine.Pin(2, machine.Pin.OUT)

connect_to_wifi()
wallet_balance = 235

while True:
    time.sleep(3)
    print(wallet_balance)