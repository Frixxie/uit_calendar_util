from setuptools import setup

setup(
    name='uit_calendar_util',
    version='0.1.0',
    description='A calendar utility aimed for UiT students',
    url='https://github.com/Frixxie/uit_calendar_util.git',
    author='Fredrik Hagen Fasteraune',
    author_email="freddahf@outlook.com",
    license='',
    packages=['uit_calendar_util'],
    install_requires=['ics',
                      'requests',
                      ],
    classifiers=[
        'Developing'
    ],
)
