from aiogram import Router, F
from aiogram.filters import CommandStart, Command
from aiogram.types import Message

router = Router()

@router.msg(Command('start'))
async def start():
    await msg.answer('иди нахуй')